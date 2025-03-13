package shuffleStyle

import (
	"math/rand"
	"theCardShuffler/cards"
	"time"
)

// ItalianShuffle simulates an Italian Shuffle on the Deck struct
func ItalianShuffle(d *cards.Deck) {
	// Split deck into two halves
	mid := len(d.Cards) / 2
	half1 := d.Cards[:mid]
	half2 := d.Cards[mid:]

	// Create a new shuffled deck
	shuffledDeck := []cards.Card{}
	rand.Seed(time.Now().UnixNano())

	// Interleave the two halves with randomness
	for len(half1) > 0 || len(half2) > 0 {
		// Randomly decide whether to pick from half1 or half2
		if len(half1) > 0 && rand.Float64() > 0.5 {
			shuffledDeck = append(shuffledDeck, half1[0])
			half1 = half1[1:]
		} else if len(half2) > 0 {
			shuffledDeck = append(shuffledDeck, half2[0])
			half2 = half2[1:]
		}
	}

	// Update the original deck with the shuffled deck
	d.Cards = shuffledDeck
}
