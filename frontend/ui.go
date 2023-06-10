package frontend

import (
	"math/rand"
)


type UI interface {
	DisplayMessage(string)
	GetAmount() int
	GetPIN() string
}

func NewUI() UI {
	return &ui{}
}

type ui struct {
	
}

func (u *ui) DisplayMessage(message string) {
	
}

func (u *ui) GetAmount() int  {
	return rand.Int()
}

func (u *ui) GetPIN() string {
	return "";
}