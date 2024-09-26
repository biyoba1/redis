package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/biyoba1/redisProject/initializer"
	"github.com/biyoba1/redisProject/internal/models"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	UserKey = "email:password"
)

func CacheUser(user models.Person) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jsonPassword, err := json.Marshal(user.Password)
	if err != nil {
		return err
	}
	err = initializer.RedisClient.HSet(ctx, UserKey, user.Email, jsonPassword).Err()
	if err != nil {
		return err
	}
	err = initializer.RedisClient.Expire(ctx, UserKey, time.Second*(day*30)).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetCacheUser(email string) (models.Person, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err := initializer.RedisClient.HGet(ctx, UserKey, email).Result()
	if err != nil {
		if err == redis.Nil {
			return models.Person{}, fmt.Errorf("Key %s not found in redis cache", email)
			//call to postgres!
		}
		return models.Person{}, fmt.Errorf("Failed to get key %s from redis cache: %w", email, err)
	}

	var user_email string
	err = json.Unmarshal([]byte(result), &user_email)
	if err != nil {
		return models.Person{}, fmt.Errorf("Failed to unmarshal data from redis cache: %w", err)
	}

	return models.Person{Email: email}, nil

}
