package initializer

import "github.com/redis/go-redis/v9"

var RedisClient *redis.Client

func ConnectToRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1234",
		DB:       0,
	})
}
