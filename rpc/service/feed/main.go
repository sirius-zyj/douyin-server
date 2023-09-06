package main

import (
	"context"
	"douyin-server/config"
	"douyin-server/database"
	"douyin-server/middleware/otel"
	feed "douyin-server/rpc/kitex_gen/feed/feedservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	config.Init()

	otel.Init(context.Background(), config.FeedOtelName)
	defer otel.Close()

	addr, _ := net.ResolveTCPAddr("tcp", config.FeedAddr)
	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.FeedServiceName}),
		server.WithRegistry(r))
	// ----------------------------

	database.Init()

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
