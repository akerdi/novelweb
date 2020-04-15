package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"novelweb/config"
)

type RedisClient struct {
	client *redis.Client
}

var RedisConnector *RedisClient

func InitRedisClient() {
	addr := fmt.Sprintf("%s:%d", config.GetConfig().Redis.Host, config.GetConfig().Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: config.GetConfig().Redis.Pass,
		DB: config.GetConfig().Redis.DB,
	})
	RedisConnector = &RedisClient{
		client: client,
	}
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
}