package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	err := redisClient.Set("test-key", "supcon", time.Hour).Err()
	if err != nil {
		log.Println(err)
		return
	}
	value, err := redisClient.Get("test-key").Result()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("test-key", value)
}

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:9295",
		Password: "supcon",
		DB:       0,
	})

	if err := redisClient.Ping().Err(); err == nil {
		log.Println("redis初始化成功")
	} else {
		panic(err)
	}

	fmt.Printf("redisClient.Ping().Val(): %v\n", redisClient.Ping().Val())
}
