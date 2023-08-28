package client

import (
	"context"
	"douyin-server/config"
	"douyin-server/rpc/kitex_gen/comment"
	"douyin-server/rpc/kitex_gen/comment/commentservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
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
	resp := new(comment.DouyinCommentActionResponse)
	resp, err := commentClient.CommentAction(context.Background(), &comment.DouyinCommentActionRequest{
		Token:       token,
		VideoId:     videoId,
		ActionType:  actionType,
		CommentText: commentText,
		CommentId:   commentId,
	})
	if err != nil {
		log.Printf("CommentAction Client get err %v\n", err)
		return nil, err
	}
	return resp, nil
}

func CommentList(token string, videoId int64) (resp *comment.DouyinCommentListResponse, err error) {
	resp = new(comment.DouyinCommentListResponse)
	resp, err = commentClient.CommentList(context.Background(), &comment.DouyinCommentListRequest{
		Token:   token,
		VideoId: videoId,
	})
	if err != nil {
		log.Printf("CommentList Client get err %v\n", err)
		return nil, err
	}
	return resp, nil
}
