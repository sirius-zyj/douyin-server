package main

import (
	"context"
	feed "douyin-server-rpc/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// ListFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) ListFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return &feed.ListFeedResponse{
		StatusCode: 200,
		VideosList: []*feed.Video{
			{
				Id:            1,
				AuthorId:      1,
				PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
				CoverUrl:      "",
				UploadTime:    "",
				Title:         "",
				FavoriteCount: 0,
				CommentCount:  0,
			},
		},
	}, nil
}

// Echo implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Echo(ctx context.Context) (resp string, err error) {
	// TODO: Your code here...
	return "this is echo", nil
}
