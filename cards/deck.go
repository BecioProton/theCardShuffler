package cards

import "fmt"

// Suit type
type Suit string

type Card struct {
	Value int
	Suit  Suit
}

type Deck struct {
	Cards []Card
}

// NewDeck creates a deck with custom suits and value range
func NewDeck(suits []Suit, minValue, maxValue int) Deck {
	deck := Deck{}
	for _, suit := range suits {
		for value := minValue; value <= maxValue; value++ {
			deck.Cards = append(deck.Cards, Card{
				Value: value,
				Suit:  suit,
			})
		}
	}
	return deck
}

// Print method for Deck
func (d Deck) Print() {
	for _, card := range d.Cards {
		fmt.Printf("%d of %s\n", card.Value, card.Suit)
	}
}
