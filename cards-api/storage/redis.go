package storage

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Ctx         context.Context
	RedisClient *redis.Client
	Timeout     time.Duration
}

func NewRedisConnection() *RedisService {
	// Connect to the Redis server
	cli := redis.NewClient(&redis.Options{
		Addr:     "cards-redis:6379",
		Password: "",
		DB:       0,
	})

	timeoutValue, err := time.ParseDuration("1h")
	if err != nil {
		log.Fatalln("unable to get timeout duration")
	}

	r := RedisService{
		Ctx:         context.Background(),
		RedisClient: cli,
		Timeout:     timeoutValue,
	}
	return &r
}

func (r *RedisService) Set(key string, value interface{}) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.RedisClient.Set(r.Ctx, key, val, r.Timeout).Err()
}

func (r *RedisService) Get(key string, dest interface{}) error {
	val, err := r.RedisClient.Get(r.Ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}
