package main

import (
	"douyin-server/config"
	"douyin-server/database"
	message "douyin-server/rpc/kitex_gen/message/messageservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	config.Init()
	addr, _ := net.ResolveTCPAddr("tcp", config.MessageAddr)
	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	svr := message.NewServer(new(MessageServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.MessageServiceName}),
		server.WithRegistry(r))
	// ----------------------------

	database.Init()

	if err = svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
