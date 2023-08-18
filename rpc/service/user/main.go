package main

import (
	"douyin-server/dao"
	user "douyin-server/rpc/kitex_gen/user/userservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8881")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := user.NewServer(new(UserServiceImpl), opts...)

	dao.Init()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
