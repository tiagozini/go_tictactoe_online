package gameroom

import (
	"sync"
	"time"
)

const (
	REFRESH_TIME = 2000
)

const (
	NO_PLAYER     = 0
	X_PLAYER      = 1
	CIRCLE_PLAYER = 2
)

const (
	PLAYER_WAITING      = 1
	PLAYER_LEAVE        = 2
	GAME_STARTED        = 3
	GAME_FINISHED       = 4
	WAITING_REMATCH     = 5
	WAITING_DESTRUCTION = 6
	WAITING_START       = 7
	GAME_DESTROYED      = 8
)

type Game struct {
	XPlayer          *User
	CirclePlayer     *User
	State            int
	Players          int
	Id               int
	Turn             int
	VitoryOf         int
	Movements        [3][3]int
	RematchDecision  *Rematch
	BothConfirmation *BothConfirmation
}

func (g *Game) Rematch() {
	var l sync.Mutex

	l.Lock()
	if g.State != WAITING_START {

		g.Movements = [3][3]int{}
		g.State = WAITING_START
		g.Turn = X_PLAYER

		users[g.XPlayer.Id] = g.CirclePlayer
		users[g.CirclePlayer.Id] = g.XPlayer

		oldX := *g.XPlayer
		*g.XPlayer = *g.CirclePlayer
		*g.CirclePlayer = oldX
		g.RematchDecision = &Rematch{Game: g.Id, XOk: false, CircleOk: false, Challenger: nil}
		g.BothConfirmation = &BothConfirmation{}
	}
	l.Unlock()
}

func (g *Game) CheckUserLogins() bool {
	if g.XPlayer != nil && !g.XPlayer.IsLogged() {
		return false
	} else if g.CirclePlayer != nil && !g.CirclePlayer.IsLogged() {
		return false
	}
	return true
}

func (g *Game) GetChallengerNames() string {
	xPlayerName, circlePlayerName := "", ""
	if g.XPlayer != nil {
		xPlayerName = g.XPlayer.Nickname
	} else {
		xPlayerName = "?"
	}
	if g.CirclePlayer != nil {
		circlePlayerName = g.CirclePlayer.Nickname
	} else {
		circlePlayerName = "?"
	}
	return (xPlayerName + "|" + circlePlayerName)
}

func (u *User) IsXPlayer() bool {
	if u.Game != nil && u.Game.XPlayer == u {
		return true
	}
	return false
}

func (u *User) PlayerLetter() string {
	if u.IsXPlayer() {
		return "X"
	} else {
		return "O"
	}
}

func (u *User) UpdateLogin() {
	t := time.Now()
	u.LastAccess = &t
}

func (u *User) IsLogged() bool {
	if (time.Since(*u.LastAccess).Nanoseconds() / 1000000) <= (REFRESH_TIME * 2) {
		return true
	}

	return false
}

var (
	users []*User // an empty list
	games []*Game // an empty list
)

func DoMove(game *Game, user *User, positionX int, positionY int) {
	var turn int
	if game.XPlayer.Id == user.Id {
		turn = X_PLAYER
	} else {
		turn = CIRCLE_PLAYER
	}
	if turn == game.Turn {
		game.Movements[positionX][positionY] = turn
		if game.Turn == X_PLAYER {
			game.Turn = CIRCLE_PLAYER
		} else {
			game.Turn = X_PLAYER
		}

		b := false

		if !b && (game.Movements[0][0] != NO_PLAYER && game.Movements[0][0] == game.Movements[0][1] && game.Movements[0][1] == game.Movements[0][2]) {
			b = true
			game.VitoryOf = game.Movements[0][0]
		}
		if !b && (game.Movements[1][0] != NO_PLAYER && game.Movements[1][0] == game.Movements[1][1] && game.Movements[1][1] == game.Movements[1][2]) {
			b = true
			game.VitoryOf = game.Movements[1][0]
		}
		if !b && (game.Movements[2][0] != NO_PLAYER && game.Movements[2][0] == game.Movements[2][1] && game.Movements[2][1] == game.Movements[2][2]) {
			b = true
			game.VitoryOf = game.Movements[2][0]
		}

		//vertical
		if !b && (game.Movements[0][0] != NO_PLAYER && game.Movements[0][0] == game.Movements[1][0] && game.Movements[1][0] == game.Movements[2][0]) {
			b = true
			game.VitoryOf = game.Movements[0][0]
		}
		if !b && (game.Movements[0][1] != NO_PLAYER && game.Movements[0][1] == game.Movements[1][1] && game.Movements[1][1] == game.Movements[2][1]) {
			b = true
			game.VitoryOf = game.Movements[0][1]
		}
		if !b && (game.Movements[0][2] != NO_PLAYER && game.Movements[0][2] == game.Movements[1][2] && game.Movements[1][2] == game.Movements[2][2]) {
			b = true
			game.VitoryOf = game.Movements[0][2]
		}

		//diagonal
		if !b && (game.Movements[0][0] != NO_PLAYER && game.Movements[0][0] == game.Movements[1][1] && game.Movements[1][1] == game.Movements[2][2]) {
			b = true
			game.VitoryOf = game.Movements[0][0]
		}
		if !b && (game.Movements[0][2] != NO_PLAYER && game.Movements[0][2] == game.Movements[1][1] && game.Movements[1][1] == game.Movements[2][0]) {
			b = true
			game.VitoryOf = game.Movements[0][2]
		}

		if b {
			game.State = GAME_FINISHED
			game.VitoryOf = turn
		}
	}
}

func AddUser(nickname string) User {
	t := time.Now()
	u := User{Nickname: nickname, Id: len(users), LastAccess: &t} // create new user
	users = append(users, &u)                                     // append user to the list of users

	var game *Game

	if len(games) > 0 && (games[len(games)-1].Players < 2 && games[len(games)-1].State != GAME_DESTROYED) {
		game = games[len(games)-1]
		game.Players = 2
		game.State = GAME_STARTED
		game.CirclePlayer = &u
		u.Game = game
	} else {
		game = &Game{XPlayer: &u, State: PLAYER_WAITING, Players: 1, Id: len(games), Turn: X_PLAYER, RematchDecision: &Rematch{Game: len(games)}}
		games = append(games, game)
		u.Game = game
	}
	return u
}

func GetUsers() []*User {
	return users
}

func GetGames() []*Game {
	return games
}
