package main

import (
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	message "douyin-server/rpc/kitex_gen/message/messageservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8886")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := message.NewServer(new(MessageServiceImpl), opts...)

	dao.Init()
	redis.InitRedis()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
