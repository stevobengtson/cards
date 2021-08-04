package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/cards/game"
	"github.com/stevobengtson/cards/storage"
)

type GameController struct {
	RedisService   *storage.RedisService
	GameRepository *game.GameRepository
}

type CreateGameInput struct {
	NumDecks    int      `json:"decks" binding:"required"`
	PlayerNames []string `json:"players" binding:"required"`
}

type DealCardsInput struct {
	Count      int    `json:"count" binding:"required"`
	PlayerUUID string `json:"player_uuid" binding:"required"`
}

type DiscardCardsInput struct {
	CardUUIDs  []string `json:"cards" binding:"required"`
	PlayerUUID string   `json:"player_uuid" binding:"required"`
}

func NewGameController(redisService *storage.RedisService, gameRepository *game.GameRepository) *GameController {
	return &GameController{
		RedisService:   redisService,
		GameRepository: gameRepository,
	}
}

// POST /game
func (g *GameController) CreateGame(c *gin.Context) {
	var f CreateGameInput
	if err := c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGame, err := game.NewGame(f.NumDecks, f.PlayerNames)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.GameRepository.SaveGame(newGame)

	c.JSON(http.StatusOK, gin.H{"data": newGame})
}

// GET /game/:key
func (g *GameController) GetGame(c *gin.Context) {
	gameObj := g.GameRepository.FindGameByKey(c.Param("key"))
	c.JSON(http.StatusOK, gin.H{"data": gameObj})
}

// PUT /game/:key/shuffle
func (g *GameController) SuffleGameDeck(c *gin.Context) {
	gameObj := g.GameRepository.FindGameByKey(c.Param("key"))
	gameObj.Deck.Shuffle()

	g.GameRepository.SaveGame(gameObj)

	c.JSON(http.StatusOK, gin.H{"data": gameObj})
}

// PUT /game/:key/deal
func (g *GameController) Deal(c *gin.Context) {
	gameObj := g.GameRepository.FindGameByKey(c.Param("key"))

	var f DealCardsInput
	if err := c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dealPlayer := gameObj.FindPlayerByUUID(f.PlayerUUID)
	myCards := gameObj.Deck.Deal(f.Count)
	dealPlayer.Deal(myCards)

	g.GameRepository.SaveGame(gameObj)

	c.JSON(http.StatusOK, gin.H{"data": gameObj})
}

// PUT /game/:key/discard
func (g *GameController) Discard(c *gin.Context) {
	gameObj := g.GameRepository.FindGameByKey(c.Param("key"))

	var f DiscardCardsInput
	if err := c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dealPlayer := gameObj.FindPlayerByUUID(f.PlayerUUID)
	discardedCards := dealPlayer.Discard(f.CardUUIDs)
	gameObj.Discard = append(gameObj.Discard, discardedCards...)

	g.GameRepository.SaveGame(gameObj)

	c.JSON(http.StatusOK, gin.H{"data": gameObj})
}

// PUT /game/:key/return
func (g *GameController) Return(c *gin.Context) {
	gameObj := g.GameRepository.FindGameByKey(c.Param("key"))

	gameObj.Return()

	g.GameRepository.SaveGame(gameObj)

	c.JSON(http.StatusOK, gin.H{"data": gameObj})
}
