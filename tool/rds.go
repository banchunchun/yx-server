package tool

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

// Set 默认设置永久的数据
func Set(client *redis.Client, key string, value interface{}) {
	//redis 对象转string
	rdsValue, _ := json.Marshal(value)
	rds := string(rdsValue)
	client.Set(key, rds, -1)
}

// Get 从redis里面获取值
func Get(client *redis.Client, key string, any interface{}) {
	sc, err := client.Get(key).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(sc), &any)
	if err != nil {
		return
	}
}
