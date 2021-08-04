package game

import (
	"errors"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/stevobengtson/cards/deck"
	"github.com/stevobengtson/cards/lib"
	"github.com/stevobengtson/cards/player"
)

type Game struct {
	UUID    string
	Key     string
	Deck    deck.Deck
	Discard []deck.Card
	Players []player.Player
}

func NewGame(numDecks int, playerNames []string) (Game, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return Game{}, errors.New("unable to generate uuid for game")
	}

	key := lib.GenerateKey()
	var players []player.Player

	for i, playerName := range playerNames {
		p, err := player.NewPlayer(playerName, i)
		if err != nil {
			return Game{}, err
		}
		players = append(players, p)
	}

	newDeck, err := deck.NewDeck(numDecks)
	if err != nil {
		return Game{}, err
	}

	return Game{
		UUID:    uuid.String(),
		Key:     key,
		Deck:    newDeck,
		Players: players,
	}, nil

}

func (g *Game) FindPlayerByUUID(uuid string) *player.Player {
	for i, v := range g.Players {
		if v.UUID == uuid {
			return &g.Players[i]
		}
	}

	return &player.Player{}
}

func (g *Game) Return() {
	g.Deck.Cards = append(g.Deck.Cards, g.Discard...)
	g.Discard = nil
}
