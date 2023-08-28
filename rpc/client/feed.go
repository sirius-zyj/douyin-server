package client

import (
	"context"
	"douyin-server/config"
	"douyin-server/rpc/kitex_gen/feed"
	"douyin-server/rpc/kitex_gen/feed/feedservice"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var feedClient feedservice.Client

func initFeedClient() {
	// 服务发现
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	c, err := feedservice.NewClient(config.FeedServiceName,
		// client.WithHostPorts(config.FeedAddr),
		client.WithResolver(r))
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
