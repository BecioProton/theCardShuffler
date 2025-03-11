package cards

// French suits as string constants
const (
	Hearts      Suit = "Hearts"
	Diamonds    Suit = "Diamonds"
	FrenchClubs Suit = "Clubs"
	Spades      Suit = "Spades"
)

// CreateFrenchDeck creates a standard 52-card French deck
func CreateFrenchDeck() Deck {
	suits := []Suit{Hearts, Diamonds, FrenchClubs, Spades}
	return NewDeck(suits, 1, 13) // Ace to King
}
