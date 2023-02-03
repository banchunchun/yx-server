package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-api/global"
	"log"
)

func LoadRedis() *redis.Client {
	url := fmt.Sprintf("%s:%s", global.CF.Redis.Host, global.CF.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Network:  global.CF.Redis.Network,
		Addr:     url,
		Password: global.CF.Redis.Pass,
		DB:       global.CF.Redis.DB,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		log.Fatalf("redis %v init error", url)
	}
	return redisClient
}
