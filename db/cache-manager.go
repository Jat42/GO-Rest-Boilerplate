package db

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

type RedisCacheManager struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCacheManager {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisCacheManager{client: client}
}

func (r *RedisCacheManager) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(key, value, expiration).Err()
}

func (r *RedisCacheManager) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisCacheManager) Delete(key string) error {
	return r.client.Del(key).Err()
}

func (r *RedisCacheManager) Close() error {
	return r.client.Close()
}
