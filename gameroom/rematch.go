package gameroom

type Rematch struct {
	Game       int
	XOk        bool
	CircleOk   bool
	Challenger *User
}

func (r *Rematch) IsOk() bool {
	return r.XOk == true && r.CircleOk == true
}

func (r *Rematch) Join(u *User) {
	if r.XOk != true && r.CircleOk != true {
		r.Challenger = u
	}
	if u.IsXPlayer() {
		r.XOk = true
	} else {
		r.CircleOk = true
	}
}

func (r *Rematch) IsChallenger(u *User) bool {
	return (r.Challenger == u)
}
