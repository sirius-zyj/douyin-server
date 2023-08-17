package client

import (
	"context"
	"douyin-server/dao"
	"douyin-server/rpc/kitex_gen/feed"
	"douyin-server/rpc/kitex_gen/feed/feedservice"
	"log"
	"time"

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

// func GetVideoByUserId(userId int64) (resp []dao.Dvideo, err error) {
// 	var respClient *feed.DouyinFeedResponse = new(feed.DouyinFeedResponse)
// 	respClient, err = feedClient.GetVideo(context.Background(), &feed.DouyinFeedRequest{
// 		LatestTime: &userId,
// 	})

// 	if err != nil {
// 		return nil, err
// 	}
// 	if respClient.StatusCode != 200 {
// 		return nil, nil
// 	}

// 	for _, v := range respClient.VideosList {
// 		resp = append(resp, dao.Dvideo{
// 			Play_url: v.PlayUrl,
// 		})
// 	}
// 	return resp, nil
// }

func GetVideoByTime(time time.Time) (resp []dao.Dvideo, err error) {
	var respClient *feed.DouyinFeedResponse = new(feed.DouyinFeedResponse)
	var latestTime int64 = time.Unix()
	respClient, err = feedClient.GetVideo(context.Background(), &feed.DouyinFeedRequest{
		LatestTime: &latestTime,
	})
	if err != nil {
		log.Printf("GetVideoByTime get err %v\n", err)
		return nil, err
	}
	if respClient.StatusCode != 0 {
		log.Printf("return StatusCode is %d\n", respClient.StatusCode)
		return nil, nil
	}

	for _, v := range respClient.VideoList {
		resp = append(resp, dao.Dvideo{
			Play_url: v.PlayUrl,
		})
	}
	return resp, nil
}
