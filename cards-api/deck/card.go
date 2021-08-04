package deck

import (
	"errors"

	uuid "github.com/nu7hatch/gouuid"
)

// Card holds the card suits and types in the deck
type Card struct {
	UUID string
	Type string
	Suit string
}

func NewCard(value string, suit string) (card Card, err error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return Card{}, errors.New("unable to generate uuid")
	}

	return Card{
		UUID: uuid.String(),
		Type: value,
		Suit: suit,
	}, nil
}
