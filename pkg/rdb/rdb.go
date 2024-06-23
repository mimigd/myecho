package rdb

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:         os.Getenv("MYECHO_REDIS"),
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     100,
		MinIdleConns: 10,
	})
	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func CloseRedis() {
	if err := RDB.Close(); err != nil {
		log.Printf("Failed to close Redis: %v", err)
	}
}
