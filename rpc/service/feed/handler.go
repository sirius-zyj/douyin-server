package main

import (
	"context"
	"douyin-server/config"
	"douyin-server/database"
	"douyin-server/database/dao"
	"douyin-server/middleware/gorse"
	"douyin-server/middleware/jwt"
	feed "douyin-server/rpc/kitex_gen/feed"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	_, span := otel.Tracer(config.FeedOtelName).Start(ctx, "Feed")
	defer span.End()

	resp = new(feed.DouyinFeedResponse) // 分配内存
	if req.Token != nil {
		userId := jwt.GetUserIdByToken(*req.Token)
		if respVideo, err := gorse.GetUserRecommend(ctx, userId); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "GetUserRecommend err")

			resp.StatusCode = 404
			resp.StatusMsg = new(string)
			*resp.StatusMsg = err.Error()
		} else {
			resp.StatusCode = 0
			for _, tmp := range respVideo {
				tmpV, _ := database.GetVideoById(tmp)
				resp.VideoList = append(resp.VideoList, database.DaoVideo2RPCVideo(req.Token, &tmpV))
			}
		}
		return
	}

	if respVideo, err := dao.GetVideoByTime(time.Unix(*req.LatestTime, 0)); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "GetVideoByTime err")

		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = err.Error()
		return resp, err
	} else {
		resp.StatusCode = 0
		for _, tmp := range respVideo {
			resp.VideoList = append(resp.VideoList, database.DaoVideo2RPCVideo(req.Token, &tmp))
		}
	}

	return resp, nil
}
