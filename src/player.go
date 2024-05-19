package src

import "fmt"

type Player struct {
	Name   string
	Cards  []Card
	Total  int
	Status string
}

func (p *Player) CalculateTotal() {
	total := 0
	for _, card := range p.Cards {
		total += card.Value
	}
	p.Total = total
}

func (p *Player) Show() {
	p.CalculateTotal()
	fmt.Printf("Player %v\n", p.Name)
	for cIdx, card := range p.Cards {
		fmt.Printf("- [%v] %v %v\n", cIdx, card.Name, card.Suit)
	}
	fmt.Printf("-- Total: %v\n", p.Total)
}

func (p *Player) ReceiveCard(card *Card) {
	p.Cards = append(p.Cards, *card)
	p.CalculateTotal()
}

func (p *Player) Play(deck *Deck) {

	for true {
		fmt.Printf("[Jogador %v] ", p.Name)
		fmt.Println("O que você quer fazer? 1 - Puxar uma carta, 2 - Parar")
		choice := ""
		fmt.Scan(&choice)

		if choice == "1" {
			p.Cards = append(p.Cards, *deck.GetCard())
			p.CalculateTotal()
			p.Show()
			if p.Total == 21 {
				fmt.Printf("%v BLACK JACK", p.Name)
				break
			}
			if p.Total > 21 {
				fmt.Printf("Estourou com %v\n", p.Total)
				break
			}
		} else if choice == "2" {
			break
		}
	}
	fmt.Printf("[Jogador %v] Finalizou a sua jogada\n", p.Name)
}
