package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/cards/controllers"
	"github.com/stevobengtson/cards/game"
	"github.com/stevobengtson/cards/storage"
)

func main() {
	redis := storage.NewRedisConnection()
	gameRepository := game.NewGameRepository(redis)

	gameCtrl := controllers.NewGameController(redis, gameRepository)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"PUT", "GET", "POST", "OPTION"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization", "Origin", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/game", gameCtrl.CreateGame)
	r.GET("/game/:key", gameCtrl.GetGame)
	r.PUT("/game/:key/shuffle", gameCtrl.SuffleGameDeck)
	r.PUT("/game/:key/deal", gameCtrl.Deal)
	r.PUT("/game/:key/discard", gameCtrl.Discard)
	r.PUT("/game/:key/return", gameCtrl.Return)

	r.Run(":8080")
}
