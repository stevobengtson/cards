

Card:
 UUID: (string)
 Type: (string) [A, 2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, Jo]
 Suit: (string) [D, H, C, S, X]

Player:
 UUID: (string)
 Name: (string)
 Position: (int)
 Cards: ([]Card)

Deck:
 Cards: ([]Card)

Game:
 UUID: (string)
 Key: (string)
 Decks: ([]Deck)
 Discard: ([]Cards)
 Players: ([]Player)

Storage Structure:
"key" : {
    "uuid": (string),
    "decks": [
      {
          "cards": [
                {
                    "uuid": (string),
                    "type": (string),
                    "suit": (string),
                },
                ...
          ]
      },
      ...
    ]
    "discard": {
        "cards": [
            {
                "uuid": (string),
                "type": (string),
                "suit": (string),
            },
            ...
        ]
    },
    "players": [
        {
            "uuid": (string),
            "name": (string),
            "position": (int),
            "cards": [
                {
                    "uuid": (string),
                    "type": (string),
                    "suit": (string),
                },
                ...
            ]
        }
    ]
}

Actions:
1. Create Game
    - Players = (json) [ { "name": (string) }, ... ]
    - Decks = (int) 1-N
1. Shuffle Decks
1. Deal Cards to Player
    - Player = (uuid)
    - Count = (int) 1-N
1. Discard Cards from Player
    - Player = (uuid)
    - Cards = (array) [ (uuid), ... ]
1. Discard to Decks

