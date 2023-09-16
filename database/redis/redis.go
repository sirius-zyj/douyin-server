package redis

import (
	"context"
	"douyin-server/config"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client //key:videoId   value:commentInfo

	Ctx           = context.Background()
	randGenerator *rand.Rand
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

	randGenerator = rand.New(rand.NewSource(time.Now().Unix()))
}

func GenExpireTime() (expireTime time.Duration) {
	// 生成随机数
	randomSeconds := randGenerator.Intn(60) // 生成一个0~59之间的随机数

	expireTime = config.Exipretime + time.Second*time.Duration(randomSeconds)
	return
}
