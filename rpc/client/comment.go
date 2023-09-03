package client

import (
	"context"
	"douyin-server/config"
	"douyin-server/rpc/kitex_gen/comment"
	"douyin-server/rpc/kitex_gen/comment/commentservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var commentClient commentservice.Client

func initCommentClient() {
	// 服务发现
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	c, err := commentservice.NewClient(config.CommentServiceName,
		// client.WithHostPorts(config.CommentAddr),
		client.WithResolver(r))

	if err != nil {
		panic(err)
	}
	commentClient = c
}

func CommentAction(token string, videoId int64, actionType int32, commentText *string, commentId *string) (*comment.DouyinCommentActionResponse, error) {
	_, span := otel.Tracer(config.RouterOtelName).Start(context.Background(), config.RouterOtelName+"-CommentAction")
	defer span.End()

	resp := new(comment.DouyinCommentActionResponse)
	resp, err := commentClient.CommentAction(context.Background(), &comment.DouyinCommentActionRequest{
		Token:       token,
		VideoId:     videoId,
		ActionType:  actionType,
		CommentText: commentText,
		CommentId:   commentId,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "CommentAction Client get err")

		return nil, err
	}
	return resp, nil
}

func CommentList(token string, videoId int64) (resp *comment.DouyinCommentListResponse, err error) {
	ctx, span := otel.Tracer(config.RouterOtelName).Start(context.Background(), config.RouterOtelName+"-CommentList")
	defer span.End()

	resp = new(comment.DouyinCommentListResponse)
	resp, err = commentClient.CommentList(ctx, &comment.DouyinCommentListRequest{
		Token:   token,
		VideoId: videoId,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "CommentList Client get err")

		return nil, err
	}
	return resp, nil
}
