package client

import (
	"context"
	"douyin-server/rpc/kitex_gen/publish"
	"douyin-server/rpc/kitex_gen/publish/publishservice"
	"log"

	"github.com/cloudwego/kitex/client"
)

var publishClient publishservice.Client

func initPublishClient() {
	c, err := publishservice.NewClient("publish", client.WithHostPorts("0.0.0.0:8884"))
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
