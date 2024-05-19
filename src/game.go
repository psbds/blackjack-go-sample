package src

import (
	"fmt"
	"strconv"
)

type Game struct {
	Players []Player
	Deck    *Deck
	Dealer  Dealer
}

func NewGame(deck *Deck) *Game {
	game := Game{
		Players: []Player{},
		Deck:    deck,
		Dealer: Dealer{
			&Player{
				Name:  "Dealer",
				Total: 0,
				Cards: []Card{},
			},
		},
	}
	fmt.Println("Digite o numero de jogadres:")
	answer := ""
	fmt.Scan(&answer)

	numberOfPlayers, err := strconv.Atoi(answer)

	if err != nil || numberOfPlayers == 0 || numberOfPlayers > 12 {
		fmt.Println("Digite apenas numeros inteiros entre 1 e 12")
		panic(err)
	}

	for i := 0; i < numberOfPlayers; i++ {
		fmt.Printf("Digite o nome do Player %v:\n", i+1)
		playerName := ""
		fmt.Scan(&playerName)
		game.Players = append(game.Players, Player{
			Name:  playerName,
			Cards: []Card{},
		})
	}

	return &game
}

func (g *Game) ShowPlayers() {
	g.Dealer.Show()
	for _, player := range g.Players {
		player.Show()
	}
}

func (g *Game) SetupGame() {
	for i := range g.Players {
		g.Players[i].ReceiveCard(g.Deck.GetCard())
		g.Players[i].ReceiveCard(g.Deck.GetCard())
	}

	g.Dealer.ReceiveCard(g.Deck.GetCard())
}

func (g *Game) StartGame() {
	fmt.Printf("O Dealer esta com %v", g.Dealer.Total)
	for i := range g.Players {
		g.Players[i].Play(g.Deck)
	}

	g.Dealer.PlayLoop(g.Deck)
}

func (g *Game) CheckWinner() {

	fmt.Printf("Total do Dealer %v\n", g.Dealer.Total)

	for _, player := range g.Players {

		if player.Total > 21 {
			player.Status = "LOSE"
		} else if g.Dealer.Total > 21 {
			player.Status = "WIN"
		} else if player.Total == g.Dealer.Total {
			player.Status = "DRAW"
		} else if player.Total > g.Dealer.Total {
			player.Status = "WIN"
		} else {
			player.Status = "LOSE"
		}
		fmt.Printf("Player %v = %v\n", player.Name, player.Status)
	}
}
