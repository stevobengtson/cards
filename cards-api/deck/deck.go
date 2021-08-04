package deck

import (
	"math/rand"
	"time"
)

// Deck holds the cards in the deck to be shuffled
type Deck struct {
	Cards []Card
}

// NewDeck creates a deck of cards to be used
func NewDeck(numDecks int) (Deck, error) {
	var deck Deck

	// Valid types include Two, Three, Four, Five, Six
	// Seven, Eight, Nine, Ten, Jack, Queen, King & Ace
	types := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	// Valid suits include Heart, Diamond, Club & Spade
	suits := []string{"H", "D", "C", "S"}

	// Loop over each type and suit appending to the deck
	for d := 0; d < numDecks; d++ {
		for i := 0; i < len(types); i++ {
			for n := 0; n < len(suits); n++ {
				card, err := NewCard(types[i], suits[n])
				if err != nil {
					return Deck{}, err
				}
				deck.Cards = append(deck.Cards, card)
			}
		}
	}

	return deck, nil
}

// Shuffle the deck
func (deck *Deck) Shuffle() {
	for i := 1; i < len(deck.Cards); i++ {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the the current card doesn't match the random
		// int we generated then we'll switch them out
		if i != r {
			deck.Cards[r], deck.Cards[i] = deck.Cards[i], deck.Cards[r]
		}
	}
}

// Deal a specified amount of cards
func (deck *Deck) Deal(n int) []Card {
	dealt := []Card{}
	for i := 0; i < n; i++ {
		dealt = append(dealt, deck.Cards[i])
		deck.Cards = append(deck.Cards[:i], deck.Cards[i+1:]...)
		if len(deck.Cards) < (n - i) {
			break
		}
	}
	return dealt
}

// Seed our randomness with the current time
func init() {
	rand.Seed(time.Now().UnixNano())
}
