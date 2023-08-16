package client

import (
	"context"
	"douyin-server/dao"
	"douyin-server/rpc/kitex_gen/feed"
	"douyin-server/rpc/kitex_gen/feed/feedservice"
	"log"

	"github.com/cloudwego/kitex/client"
)

var feedClient feedservice.Client

func initFeedClient() {
	c, err := feedservice.NewClient("feed", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	feedClient = c
}

func GetVideoByUserId(userId int64) (resp []dao.Dvideo, err error) {
	var respClient *feed.FeedResponse
	respClient, err = feedClient.GetVideo(context.Background(), &feed.FeedRequest{
		AuthorId: &userId,
	})

	if err != nil || respClient.StatusCode != 200 {
		return nil, err
	}

	log.Println(resp)

	return nil, nil
}

func Test() {
	c, err := feedservice.NewClient("feed", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	feedClient = c

	// resp, err := feedClient.Echo(context.Background())
	// resp, err := feedClient.ListFeed(context.Background(), &feed.FeedRequest{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp)

}
