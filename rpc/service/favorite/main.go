package main

import (
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	favorite "douyin-server/rpc/kitex_gen/favorite/favoriteservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8882")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := favorite.NewServer(new(FavoriteServiceImpl), opts...)

	dao.Init()
	redis.InitRedis()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
