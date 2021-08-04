package player

import (
	"errors"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/stevobengtson/cards/deck"
)

type Player struct {
	UUID     string
	Name     string
	Position int
	Cards    []deck.Card
}

func NewPlayer(name string, position int) (player Player, err error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return Player{}, errors.New("unable to generate uuid for player")
	}

	return Player{
		UUID:     uuid.String(),
		Name:     name,
		Position: position,
		Cards:    []deck.Card{},
	}, nil
}

func (p *Player) Deal(cards []deck.Card) {
	p.Cards = append(p.Cards, cards...)
}

func (p *Player) FindAndRemoveCardByUUID(uuid string) deck.Card {
	for i, v := range p.Cards {
		if v.UUID == uuid {
			p.Cards = append(p.Cards[:i], p.Cards[i+1:]...)
			return v
		}
	}

	return deck.Card{}
}

func (p *Player) Discard(cardUUIDs []string) []deck.Card {
	discarded := []deck.Card{}
	for _, uuid := range cardUUIDs {
		discarded = append(discarded, p.FindAndRemoveCardByUUID(uuid))
	}
	return discarded
}
