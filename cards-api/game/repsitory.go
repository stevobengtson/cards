package game

import "github.com/stevobengtson/cards/storage"

type GameRepository struct {
	redisService *storage.RedisService
}

func NewGameRepository(redisService *storage.RedisService) *GameRepository {
	return &GameRepository{redisService: redisService}
}

func (r *GameRepository) FindGameByKey(key string) Game {
	var gameObj Game
	err := r.redisService.Get(key, &gameObj)
	if err != nil {
		return Game{}
	}

	return gameObj
}

func (r *GameRepository) SaveGame(gameObj Game) error {
	return r.redisService.Set(gameObj.Key, gameObj)
}
