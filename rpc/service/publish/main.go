package main

import (
	"douyin-server/dao"
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

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
