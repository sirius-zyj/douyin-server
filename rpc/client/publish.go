package client

import (
	"context"
	"douyin-server/config"
	"douyin-server/rpc/kitex_gen/publish"
	"douyin-server/rpc/kitex_gen/publish/publishservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient publishservice.Client

func initPublishClient() {
	// 服务发现
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	c, err := publishservice.NewClient(config.PublishServiceName,
		// client.WithHostPorts(config.PublishAddr),
		client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	publishClient = c
}

func Publish(token string, videoData []byte, title string) (resp *publish.DouyinPublishActionResponse, err error) {
	resp = new(publish.DouyinPublishActionResponse)
	resp, err = publishClient.Publish(context.Background(), &publish.DouyinPublishActionRequest{
		Token: token,
		Data:  videoData,
		Title: title,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func PublishList(userId int64, token string) (resp *publish.DouyinPublishListResponse, err error) {
	resp = new(publish.DouyinPublishListResponse)
	resp, err = publishClient.PublishList(context.Background(), &publish.DouyinPublishListRequest{
		UserId: userId,
		Token:  token,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
