package cache

import (
	"github.com/go-redis/redis"
)

var RedisClient redis.UniversalClient

func NewRedisClient(password string, redisType int, addrs []string, host string, port string, db int, poolSize int) {

	opt := &redis.UniversalOptions{
		Password: password,
	}

	switch redisType {
	case 2: // 集群
		opt.Addrs = addrs
		if len(opt.Addrs) < 2 {
			panic("集群地址数量错误")
		}
	default: // 单点
		opt.Addrs = []string{host + ":" + port}
		opt.DB = db
	}

	if poolSize > 0 {
		opt.PoolSize = poolSize
	}

	RedisClient = redis.NewUniversalClient(opt)

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("连接 redis 失败：" + err.Error())
	}
}
