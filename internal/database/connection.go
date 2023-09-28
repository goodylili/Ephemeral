package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	Client *redis.Client
}

func NewRedisInstance() (*redis.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf("host=%v port=%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	options, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(options)

	return redisClient, nil
}
