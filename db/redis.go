package db

import (
	"github.com/go-redis/redis"
	. "video/common"
	"video/logger"
)

var client *redis.Client

func GetRedisClient() error {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     REDIS_ADDR,
			Password: REDIS_PASSWORD,
			DB:       REDIS_database,
		})
	}
	_, err := client.Ping().Result()
	if err != nil {
		logger.Error("fail to connect redis, error:", err)
		return err
	}
	logger.Info("succeed connect to redis")
	return nil
}


func GetClient() *redis.Client {
	if client == nil {
		GetRedisClient()
	}
	return client
}


