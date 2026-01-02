package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitRedis () {
	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv(""),
		Password: "",
		DB: 0,})
}