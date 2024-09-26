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
	day = 86400
	key = "products:price"
)

func CacheProduct(product models.Product) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jsonPrice, err := json.Marshal(product.Price)
	if err != nil {
		return err
	}

	err = initializer.RedisClient.HSet(ctx, key, product.Name, jsonPrice).Err()
	if err != nil {
		return err
	}

	err = initializer.RedisClient.Expire(ctx, key, time.Second*(day*30)).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetCacheProduct(name string) (models.Product, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err := initializer.RedisClient.HGet(ctx, key, name).Result()

	if err != nil {
		if err == redis.Nil {
			return models.Product{}, fmt.Errorf("Key %s not found in redis cache", name)
		}
		return models.Product{}, fmt.Errorf("Failed to get key %s from redis cache: %w", name, err)
	}

	var price float64
	err = json.Unmarshal([]byte(result), &price)
	if err != nil {
		return models.Product{}, fmt.Errorf("Failed to unmarshal data from redis cache: %w", err)
	}

	return models.Product{Name: name, Price: price}, nil
}
