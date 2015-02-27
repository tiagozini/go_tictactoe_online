package gameroom

import "time"

type User struct {
	Nickname   string
	Id         int
	Game       *Game
	LastAccess *time.Time
}
