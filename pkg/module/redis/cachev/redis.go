package cachev

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func InitRedis(addr, password string, db int) error {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	return nil
}
