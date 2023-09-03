package main

import (
	"context"
	"douyin-server/config"
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	"douyin-server/middleware/otel"
	comment "douyin-server/rpc/kitex_gen/comment/commentservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	config.Init()

	otel.Init(context.Background(), config.CommentOtelName)
	defer otel.Close()

	addr, _ := net.ResolveTCPAddr("tcp", config.CommentAddr)
	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.CommentServiceName}),
		server.WithRegistry(r))
	// ----------------------------

	dao.Init()
	redis.InitRedis()

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
