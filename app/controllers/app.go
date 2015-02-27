package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"go_tictactoe_online/gameroom"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Welcome to Tic-Tac-Toe Combate"
	return c.Render(greeting)
}

func (c App) Welcome() revel.Result {
	greeting := "Welcome my place"
	id := c.Params.Get("id")
	return c.Render(greeting, id)
}

func (c App) EnterGame(nickname string) revel.Result {
	c.Validation.Required(nickname).Message("Your nickname is required!")
	c.Validation.MinSize(nickname, 3).Message("Your nickname is not long enough!")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	user := gameroom.AddUser(nickname)
	return c.Redirect(fmt.Sprintf("/Refresh/Room?userId=%d", user.Id))
}
