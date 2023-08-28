package main

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	publish "douyin-server/rpc/kitex_gen/publish/publishservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", config.PublishAddr)
	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.PublishServiceName}),
		server.WithRegistry(r))
	// ----------------------------

	dao.Init()
	dao.Oss_init()
	redis.InitRedis()

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
