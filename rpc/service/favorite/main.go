package main

import (
	"douyin-server/config"
	"douyin-server/database"
	favorite "douyin-server/rpc/kitex_gen/favorite/favoriteservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	config.Init()
	addr, _ := net.ResolveTCPAddr("tcp", config.FavoriteAddr)
	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.FavoriteServiceName}),
		server.WithRegistry(r))
	// ----------------------------

	database.Init()

	if err = svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
