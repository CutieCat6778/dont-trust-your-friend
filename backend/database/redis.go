package database

import (
	"context"
	"cutiecat6778/dont-trust-your-friend/lib"

	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	redis.UniversalClient
}

func NewRedisHandler() (*RedisHandler, *lib.CustomError) {
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{"localhost:6379"},
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, lib.NewError("Redis Connection Error", 500, lib.RedisService)
	}

	return &RedisHandler{redisClient}, nil
}

func (r *RedisHandler) GetValue(key string) (string, *lib.CustomError) {
	val, err := r.Get(context.Background(), key).Result()
	if err != nil {
		return "", lib.NewError("Redis Get Error", 500, lib.RedisService)
	}

	return val, nil
}

func (r *RedisHandler) SetValue(key string, value string) *lib.CustomError {
	err := r.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return lib.NewError("Redis Set Error", 500, lib.RedisService)
	}

	return nil
}

func (r *RedisHandler) DeleteValue(key string) *lib.CustomError {
	err := r.Del(context.Background(), key).Err()
	if err != nil {
		return lib.NewError("Redis Delete Error", 500, lib.RedisService)
	}

	return nil
}
