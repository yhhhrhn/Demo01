package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func NewClient(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "AnotherWine1130",
		DB:       0,
	})

	// Perform basic diagnostic to check if the connection is working
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}

	//fmt.Println(pong)
	return rdb
}
