package main

import (
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	publish "douyin-server/rpc/kitex_gen/publish/publishservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8884")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := publish.NewServer(new(PublishServiceImpl), opts...)

	dao.Init()
	dao.Oss_init()
	redis.InitRedis()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
