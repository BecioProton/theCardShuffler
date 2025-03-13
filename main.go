package main

import (
	"fmt"
	"theCardShuffler/cards"
	"theCardShuffler/shuffleStyle"
)

func main() {
	fmt.Println("French Deck:")
	frenchDeck := cards.CreateFrenchDeck()
	frenchDeck.Print()

	shuffleStyle.ItalianShuffle(frenchDeck)

	fmt.Println("\nItalian Deck:")
	italianDeck := cards.CreateItalianDeck()
	italianDeck.Print()

}
