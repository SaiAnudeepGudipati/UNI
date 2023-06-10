package main

import (
	"UNI/authorisation"
	"UNI/cashDispenser"
	"UNI/frontend"
)

func main() {
	c := make(chan int)
	
	cardHolder := frontend.NewCardHolder()
	ui := frontend.NewUI()
	
	aServer := authorisation.NewAuthServer()
	
	dispenser := cashDispenser.NewDispenser()
	
	
	
	err := cardHolder.RegisterCallbackChan(c)
	if(err != nil){
		//
		return
	}

	for true {
		select {
		case <-c:
			err = cardHolder.Lock()
			if(err != nil){
				ui.DisplayMessage("Collect Card, Unable to process transactions")

				return
			}

			cardDetails := cardHolder.GetCardDetails()

			amt := ui.GetAmount()

			pin := ui.GetPIN()

			isSuccess, err := aServer.GetAuth(cardDetails, amt, pin)

			if(err != nil){
				err = cardHolder.Unlock()
				if(err != nil){
					ui.DisplayMessage("Contact Help")

					return
				}
				ui.DisplayMessage("Collect Card, Unable to authorise transactions")

				return
			}

			if isSuccess{
				err = dispenser.AddCash(amt)
				if err != nil {
					ui.DisplayMessage("Unable to dispense cash, reverting transaction")

					return
				}

				err = dispenser.DispenseCash()
				if err != nil {
					ui.DisplayMessage("Unable to dispense cash, reverting transaction")

					return
				}

				err = cardHolder.Unlock()
				if(err != nil){
					ui.DisplayMessage("Contact Help")

					return
				}
			} else {

				err = cardHolder.Unlock()
				if(err != nil){
					ui.DisplayMessage("Contact Help")

					return
				}

				ui.DisplayMessage("Collect Card, Authorisation failed")
			}

			//	case 
		}
	}
}