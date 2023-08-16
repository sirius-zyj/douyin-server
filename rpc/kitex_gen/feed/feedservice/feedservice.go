// Code generated by Kitex v0.7.0. DO NOT EDIT.

package feedservice

import (
	"context"
	feed "douyin-server/rpc/kitex_gen/feed"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return feedServiceServiceInfo
}

var feedServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FeedService"
	handlerType := (*feed.FeedService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetVideo": kitex.NewMethodInfo(getVideoHandler, newFeedServiceGetVideoArgs, newFeedServiceGetVideoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "feed",
		"ServiceFilePath": "idl/feed.thrift",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.0",
		Extra:           extra,
	}
	return svcInfo
}

func getVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*feed.FeedServiceGetVideoArgs)
	realResult := result.(*feed.FeedServiceGetVideoResult)
	success, err := handler.(feed.FeedService).GetVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedServiceGetVideoArgs() interface{} {
	return feed.NewFeedServiceGetVideoArgs()
}

func newFeedServiceGetVideoResult() interface{} {
	return feed.NewFeedServiceGetVideoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetVideo(ctx context.Context, req *feed.FeedRequest) (r *feed.FeedResponse, err error) {
	var _args feed.FeedServiceGetVideoArgs
	_args.Req = req
	var _result feed.FeedServiceGetVideoResult
	if err = p.c.Call(ctx, "GetVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
