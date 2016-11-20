package main

import (
	"fmt"
	"gopkg.in/redis.v5"
)

//define redis client
var redisClient *redis.Client

//initialize redis client
func initRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     *redisAddr,
		Password: *redisPW,
		DB:       *redisDB,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		return fmt.Errorf("Error while ping redis server: %s ", err)
	}
	return nil
}

//write data to redis
func redisSet(key, value string) error {
	err := redisClient.Set(key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("error while redis set: %s", err)
	}
	return nil
}

//read data from redis
func redisGet(key string) (string, error) {
	value, err := redisClient.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("error while redis get: %s", err)
	}
	return string(value), nil
}
