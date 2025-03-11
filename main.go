package main

import (
	"fmt"
	"theCardShuffler/cards"
)

func main() {
	fmt.Println("French Deck:")
	frenchDeck := cards.CreateFrenchDeck()
	frenchDeck.Print()

	fmt.Println("\nItalian Deck:")
	italianDeck := cards.CreateItalianDeck()
	italianDeck.Print()
}
