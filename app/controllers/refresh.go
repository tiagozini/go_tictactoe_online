package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/tiagozini/go_tictactoe_online/gameroom"
)

type Refresh struct {
	*revel.Controller
}

func (c Refresh) Room(userId int) revel.Result {

	users := gameroom.GetUsers()
	user := users[userId]
	game := user.Game

	var action string
	var message string
	refresh := "0"

	if !game.CheckUserLogins() {
		game.State = gameroom.PLAYER_LEAVE
	}
	switch {
	case game.State == gameroom.GAME_FINISHED:
		message = "Game Finished!"
		action = "finished"
	case game.State == gameroom.PLAYER_WAITING ||
		game.State == gameroom.GAME_STARTED:

		if game.Players == 2 {
			if (user.PlayerLetter() == "X" && game.Turn == gameroom.X_PLAYER) || (user.PlayerLetter() == "O" && game.Turn == gameroom.CIRCLE_PLAYER) {
				message = "Play!!!!"
				action = "playing"
			} else {
				message = "Wait your mate do his move..."
				action = "waiting"
			}
		} else if game.Players == 1 {
			message = "Wait the other player enter..."
			action = "waiting"
		}
	case game.State == gameroom.WAITING_REMATCH:

		if game.RematchDecision.IsChallenger(user) {
			message = "Wait mate answer..."
			action = "waiting_rematch"
		} else {
			message = "Decide if you accept the challenge."
			action = "waiting_rematch_answer"
		}
	case game.State == gameroom.WAITING_START:

		if game.BothConfirmation.IsOk() {
			game.State = gameroom.GAME_STARTED
			return c.Redirect(fmt.Sprintf("Room?userId=%d", userId))
		} else {
			message = "Wait the another mate be right..."
			action = "waiting"
			refresh = "1"
		}
		game.BothConfirmation.Join(user)
	case game.State == gameroom.WAITING_DESTRUCTION:

		message = "The game terminated. The challenger was declined."
		action = "waiting_destruction"
		game.State = gameroom.GAME_DESTROYED
	case game.State == gameroom.PLAYER_LEAVE:

		message = "The mate leave the room. Try to play a new game."
		action = "waiting_destruction"
		game.State = gameroom.GAME_DESTROYED
	case game.State == gameroom.GAME_DESTROYED:

		c.Redirect(App.Index)
	default:

		c.Redirect(App.Index)
	}
	playerLetter := user.PlayerLetter()
	return c.Render(user, game, playerLetter, message, action, refresh)
}

func (c Refresh) Play(userId int, x int, y int) revel.Result {
	users := gameroom.GetUsers()
	user := users[userId]
	game := user.Game

	gameroom.DoMove(game, user, x, y)
	return c.Redirect(fmt.Sprintf("Room?userId=%d", userId))
}

func (c Refresh) AskRematch(userId int) revel.Result {
	users := gameroom.GetUsers()
	user := users[userId]
	game := user.Game
	game.RematchDecision.Join(user)
	game.State = gameroom.WAITING_REMATCH
	if game.RematchDecision.IsOk() {
		game.Rematch()
	}
	return c.Redirect(fmt.Sprintf("Room?userId=%d", userId))
}

func (c Refresh) RejectRematch(userId int) revel.Result {
	users := gameroom.GetUsers()
	user := users[userId]
	user.Game.State = gameroom.WAITING_DESTRUCTION

	return c.Redirect(fmt.Sprintf("/App/Index"))
}

func (c Refresh) RefreshSession(userId int) revel.Result {
	users := gameroom.GetUsers()
	user := users[userId]
	user.UpdateLogin()
	return c.RenderHtml(``)
}
