package main

import (
	"context"
	"douyin-server/dao"
	feed "douyin-server/rpc/kitex_gen/feed"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	resp = new(feed.DouyinFeedResponse) // 分配内存
	if respVideo, err := dao.GetVideoByTime(time.Unix(*req.LatestTime, 0)); err != nil {
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = err.Error()
		return resp, err
	} else {
		resp.StatusCode = 0
		for _, tmp := range respVideo {
			resp.VideoList = append(resp.VideoList, dao.DaoVideo2RPCVideo(req.Token, &tmp))
		}
	}

	return resp, nil
}
