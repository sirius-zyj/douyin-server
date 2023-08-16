package client

import (
	"context"
	"douyin-server-rpc/kitex_gen/feed/feedservice"
	"log"

	"github.com/cloudwego/kitex/client"
)

var feedClient feedservice.Client

func Test() {
	c, err := feedservice.NewClient("feed", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	feedClient = c

	resp, err := feedClient.Echo(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

}
