package client

import (
	"context"
	"douyin-server/rpc/kitex_gen/feed"
	"douyin-server/rpc/kitex_gen/feed/feedservice"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
)

var feedClient feedservice.Client

func initFeedClient() {
	c, err := feedservice.NewClient("feed", client.WithHostPorts("0.0.0.0:8880"))
	if err != nil {
		log.Fatal(err)
	}
	feedClient = c
}

func Feed(time time.Time, token *string) (resp *feed.DouyinFeedResponse, err error) {
	resp = new(feed.DouyinFeedResponse)
	latestTime := time.Unix()
	resp, err = feedClient.Feed(context.Background(), &feed.DouyinFeedRequest{
		LatestTime: &latestTime,
		Token:      token,
	})
	if err != nil {
		log.Printf("FeedClient get err %v\n", err)
		return nil, err
	}
	return resp, nil
}
