package main

import (
	"mysql2redis/mysql2redis"
	"time"
)

func main() {
	servicer := mysql2redis.Mysql2RedisService{}
	servicer.Service.Init(&servicer, nil, nil, map[string]interface{}{
		"ConfigPath": "./config.toml",
	})
	servicer.OnInit()
	for true {
		time.Sleep(10 * time.Second)
	}
}
