package redis

import (
	"context"
	"douyin-server/config"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client //key:videoId   value:commentInfo

	Ctx = context.Background()
)

func InitRedis() {
	addr := config.RedisAddr
	password := config.Password
	poolSize := config.PoolSize
	minConns := config.MinConns
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minConns,
	})
}
