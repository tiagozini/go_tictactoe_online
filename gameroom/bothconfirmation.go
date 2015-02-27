package gameroom

type BothConfirmation struct {
	First    *User
	FirstOk  bool
	SecondOk bool
}

func (b *BothConfirmation) Join(u *User) {
	if b.First == nil {
		b.First = u
		b.FirstOk = true
	} else if b.First != u {
		b.SecondOk = true
	}
}

func (b *BothConfirmation) IsOk() bool {
	return (b.FirstOk == true && b.SecondOk == true)
}
