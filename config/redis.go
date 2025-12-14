package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	LoadEnv()
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	addr := redisHost + ":" + redisPort

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // add if you set one
		DB:       0,
	})
}