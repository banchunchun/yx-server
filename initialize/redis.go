package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-api/global"
	"log"
)

func LoadRedis() *redis.Client {
	url := fmt.Sprintf("%s:%s", global.Config.Redis.Host, global.Config.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Network:  global.Config.Redis.Network,
		Addr:     url,
		Password: global.Config.Redis.Pass,
		DB:       global.Config.Redis.DB,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		log.Fatalf("redis %v init error", url)
	}
	fmt.Printf("init redis  %v success", url)
	return redisClient
}
