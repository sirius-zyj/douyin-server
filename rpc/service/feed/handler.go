package main

import (
	"context"
	"douyin-server/dao"
	feed "douyin-server/rpc/kitex_gen/feed"
	"log"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetVideo implements the FeedServiceImpl interface.
// GetVideo implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideo(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	resp = new(feed.DouyinFeedResponse) // 分配内存
	if req.LatestTime != nil {
		respVideo, err := dao.GetVideoByTime(time.Unix(*req.LatestTime, 0))
		if err != nil {
			resp.StatusCode = 404
			return resp, err
		}
		resp.StatusCode = 0
		for _, v := range respVideo {
			resp.VideoList = append(resp.VideoList, &feed.Video{
				PlayUrl: v.Play_url,
			})
		}
	}
	log.Printf("resp is %v\n", resp)
	return resp, nil
}
