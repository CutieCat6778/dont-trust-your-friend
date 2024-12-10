package handlers

import (
	"cutiecat6778/dont-trust-your-friend/database"
	"cutiecat6778/dont-trust-your-friend/lib"
)

type Redis struct {
	database.RedisHandler
}

func NewRedis() (*Redis, *lib.CustomError) {
	redisHandler, err := database.NewRedisHandler()
	if err != nil {
		return nil, err
	}

	return &Redis{*redisHandler}, nil
}
