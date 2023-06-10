package authorisation

import "math/rand"

func NewAuthServer() AuthServer {
	return &aServer{}
}

type AuthServer interface{
	GetAuth(string, int, string) (bool, error)
	
	
//	Reversal()
}

type aServer struct {
	
}

func (a *aServer) GetAuth(cardDetails string, amt int, pin string) (bool, error) {
	return rand.Int()%100 != 0, nil
}