package main

import (
	"douyin-server/dao"
	comment "douyin-server/rpc/kitex_gen/comment/commentservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8883")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := comment.NewServer(new(CommentServiceImpl), opts...)

	dao.Init()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
