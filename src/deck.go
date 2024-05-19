package src

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Suit   string
	Name   string
	Value  int
	Picked bool
}

func (d *Deck) Show() string {
	str := "---------------------------------------------\n"
	for idx, card := range d.Cards {

		if idx != 0 && idx%13 == 0 {
			str += "\n"
		}
		str += fmt.Sprintf("| %v %v ", card.Name, card.Suit)

	}
	str += "\n---------------------------------------------"
	return str
}

func CreateDeck() Deck {
	deck := Deck{
		Cards: []Card{},
	}
	suits := []string{"♠", "♥", "♧", "♢"}
	cards := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for _, card := range cards {
			deck.Cards = append(deck.Cards, Card{
				Suit:  suit,
				Name:  card,
				Value: getValue(card),
			})
		}
	}

	return deck
}

func (d *Deck) Flush() {
	numberOfFlushes := randInt(100, 500)
	fmt.Printf("Embaralhando %v vezes\n", numberOfFlushes)
	for i := 0; i < numberOfFlushes; i++ {
		n1 := randInt(0, 51)
		n2 := randInt(0, 51)

		card1 := d.Cards[n1]
		card2 := d.Cards[n2]

		d.Cards[n1] = card2
		d.Cards[n2] = card1

	}
}

func (d *Deck) GetCard() *Card {
	ok := false
	index := randInt(0, 51)

	for !ok {
		if index == 52 {
			index = 0
		}
		if d.Cards[index].Picked {
			index = index + 1
			continue
		}
		d.Cards[index].Picked = true
		fmt.Printf("Puxou a carta '%v %v'\n", d.Cards[index].Name, d.Cards[index].Suit)
		return &d.Cards[index]
	}

	return nil
}

func getValue(card string) int {
	if card == "A" {
		return 1
	} else if card == "J" || card == "Q" || card == "K" {
		return 10
	} else {
		value, _ := strconv.Atoi(card)
		return value
	}
}

func randInt(min, max int) int {
	return min + rand.Intn((max - min + 1))
}
