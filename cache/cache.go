package cache

import (
	"context"
	"encoding/json"
	"log"

	"2_TaskManager/config"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		panic("Unable to connect to Redis: " + err.Error())
	}

	log.Println("Connected to Redis")
}

func GetCachedTasks(key string, target interface{}) bool {
	data, err := Rdb.Get(Ctx, key).Result()
	if err != nil {
		return false
	}

	err = json.Unmarshal([]byte(data), target)
	return err == nil
}

func SetCachedTasks(key string, value interface{}) {
	taskJSON, _ := json.Marshal(value)
	Rdb.Set(Ctx, key, taskJSON, config.TaskCacheTTL)
}

func DeleteCache(key string) {
	Rdb.Del(Ctx, key)
}
