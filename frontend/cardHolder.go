package frontend

import "errors"

type CardHolder interface {
	RegisterCallbackChan(chan int) error
	GetCardDetails() string
	Lock() error
	Unlock() error
}

func NewCardHolder() CardHolder{
	return &cardHolder{}
}

type cardHolder struct {
	ch chan int
	lock bool
}

func (c *cardHolder) GetCardDetails() string {
	return ""
}

func (c *cardHolder) RegisterCallbackChan(ch chan int) error {
	c.ch = ch
	return nil
}

func (c *cardHolder) Lock() error {
	if c.lock{
		return errors.New("already Locked")
	}
	c.lock = true
	return nil
}

func (c *cardHolder) Unlock() error {
	if !c.lock{
		return errors.New("already Unlocked")
	}
	c.lock = false
	return nil
}