package main

import (
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	feed "douyin-server/rpc/kitex_gen/feed/feedservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8880")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := feed.NewServer(new(FeedServiceImpl), opts...)

	dao.Init()
	redis.InitRedis()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
