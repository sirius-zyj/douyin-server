package main

import (
	"context"
	"douyin-server/dao"
	feed "douyin-server/rpc/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetVideo implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideo(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	if *req.AuthorId != 0 {
		respVideo, err := dao.GetVideoByUserId(*req.AuthorId)
		if err != nil {
			resp.StatusCode = 404
			return resp, err
		}
		resp.StatusCode = 200
		for _, v := range respVideo {
			resp.VideosList = append(resp.VideosList, &feed.Video{
				PlayUrl: v.Play_url,
			})
		}
	}
	return resp, nil
}
