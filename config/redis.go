package config

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	LoadEnv()
	// redisHost := os.Getenv("REDIS_HOST")
	// redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	// redisUrl := os.Getenv("REDIS_URL")
	// log.Printf("DEBUG redisulr %v" , redisUrl)
	// log.Printf("DEBUG redisPassword %v" , redisPassword)
	redisHost:= os.Getenv("REDISHOST")
	redisPort:= os.Getenv("REDISPORT")
	log.Printf("DEBUG redisHost %v" , redisHost)
	log.Printf("DEBUG redisPort %v" , redisPort)
	// addr := redisHost + ":" + redisPort

	return redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword, // add if you set one
		DB:       0,
		
	})
}