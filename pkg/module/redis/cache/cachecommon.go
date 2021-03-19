package cache

import (
	"time"

	"github.com/go-redis/redis"
)

func Put(key string, value interface{}, exp time.Duration) {
	result := RedisClient.Set(key, value, exp)
	if result.Err() != nil {
		// log fmt.Sprintf("%+v", result)
	}
}

func Get(key string) string {
	val, err := RedisClient.Get(key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		// log key, val, err
		return ""
	}
	return val
}
