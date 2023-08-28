package main

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	user "douyin-server/rpc/kitex_gen/user/userservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", config.UserAddr)
	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.UserServiceName}),
		server.WithRegistry(r))
	// ----------------------------

	dao.Init()
	redis.InitRedis()

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
