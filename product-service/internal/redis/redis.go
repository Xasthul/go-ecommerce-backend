package redisdb

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(address string, password string) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
		return nil, err
	}

	return redisClient, nil
}
