package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var (
	RedisClientInstance *redis.Client
	redisOnce           sync.Once
)

func RedisInstance(db int) *redis.Client {
	if RedisClientInstance == nil {
		redisOnce.Do(func() {
			RedisClientInstance = redis.NewClient(&redis.Options{
				Addr:       "redis:6379", //os.Getenv("REDIS_HOST")
				Password:   "",
				MaxRetries: 3,
			})
		})
	}

	_, err := RedisClientInstance.Ping().Result()

	if err != nil {
		fmt.Errorf("redis connect ping failed, err: %v\n", err.Error())
	}

	RedisClientInstance.Do("select", db)
	return RedisClientInstance
}
