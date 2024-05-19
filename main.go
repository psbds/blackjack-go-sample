package main

import (
	"fmt"

	"github.com/psbds/blackjack/src"
)

func main() {

	// Cria o Deck
	deck := src.CreateDeck()
	fmt.Println(deck.Show())

	// Embaralha
	deck.Flush()
	fmt.Println(deck.Show())

	// Cria o Jogo
	game := src.NewGame(&deck)
	game.ShowPlayers()

	// Faz o Setup do Jogo
	game.SetupGame()
	game.ShowPlayers()

	// Starta o jogo
	game.StartGame()

	// Verifica vencedores
	game.CheckWinner()
}
