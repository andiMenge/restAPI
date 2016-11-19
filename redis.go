package main

import (
	"fmt"
	"gopkg.in/redis.v5"
)

const (
	redisAddr = "localhost:6379"
	redisPW   = "" // no password set
	redisDB   = 0  // use default DB
)

func initRedis() error {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPW,
		DB:       redisDB,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		return fmt.Errorf("Error while ping redis server: %s ", err)
	}
	return nil
}
