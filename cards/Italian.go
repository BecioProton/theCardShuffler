package cards

// Italian suits as string constants
const (
	Cups         Suit = "Cups"
	Coins        Suit = "Coins"
	Swords       Suit = "Swords"
	ItalianClubs Suit = "Clubs"
)

// CreateItalianDeck creates a standard 40-card Italian deck
func CreateItalianDeck() Deck {
	suits := []Suit{Cups, Coins, Swords, ItalianClubs}
	return NewDeck(suits, 1, 10) // 1 to 10
}
