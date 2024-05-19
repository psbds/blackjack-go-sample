package src

import "fmt"

type Dealer struct {
	*Player
}

func (d *Dealer) PlayLoop(deck *Deck) {
	fmt.Println("Dealer esta jogando...")
	for d.Total < 17 {
		d.ReceiveCard(deck.GetCard())
	}
}
