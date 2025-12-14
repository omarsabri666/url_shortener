package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	LoadEnv()
	// redisHost := os.Getenv("REDIS_HOST")
	// redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisUrl := os.Getenv("REDIS_URL")
	// addr := redisHost + ":" + redisPort

	return redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword, // add if you set one
		DB:       0,
		
	})
}