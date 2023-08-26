package main

import (
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	relation "douyin-server/rpc/kitex_gen/relation/relationservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8885")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := relation.NewServer(new(RelationServiceImpl), opts...)

	dao.Init()
	redis.InitRedis()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
