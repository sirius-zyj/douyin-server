// Code generated by Kitex v0.7.0. DO NOT EDIT.

package relationservice

import (
	"context"
	relation "douyin-server/rpc/kitex_gen/relation"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"RelationAction":       kitex.NewMethodInfo(relationActionHandler, newRelationServiceRelationActionArgs, newRelationServiceRelationActionResult, false),
		"RelationFollowList":   kitex.NewMethodInfo(relationFollowListHandler, newRelationServiceRelationFollowListArgs, newRelationServiceRelationFollowListResult, false),
		"RelationFollowerList": kitex.NewMethodInfo(relationFollowerListHandler, newRelationServiceRelationFollowerListArgs, newRelationServiceRelationFollowerListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "relation",
		"ServiceFilePath": "idl/relation.thrift",
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

func relationActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceRelationActionArgs)
	realResult := result.(*relation.RelationServiceRelationActionResult)
	success, err := handler.(relation.RelationService).RelationAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceRelationActionArgs() interface{} {
	return relation.NewRelationServiceRelationActionArgs()
}

func newRelationServiceRelationActionResult() interface{} {
	return relation.NewRelationServiceRelationActionResult()
}

func relationFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceRelationFollowListArgs)
	realResult := result.(*relation.RelationServiceRelationFollowListResult)
	success, err := handler.(relation.RelationService).RelationFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceRelationFollowListArgs() interface{} {
	return relation.NewRelationServiceRelationFollowListArgs()
}

func newRelationServiceRelationFollowListResult() interface{} {
	return relation.NewRelationServiceRelationFollowListResult()
}

func relationFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceRelationFollowerListArgs)
	realResult := result.(*relation.RelationServiceRelationFollowerListResult)
	success, err := handler.(relation.RelationService).RelationFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceRelationFollowerListArgs() interface{} {
	return relation.NewRelationServiceRelationFollowerListArgs()
}

func newRelationServiceRelationFollowerListResult() interface{} {
	return relation.NewRelationServiceRelationFollowerListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (r *relation.DouyinRelationActionResponse, err error) {
	var _args relation.RelationServiceRelationActionArgs
	_args.Req = req
	var _result relation.RelationServiceRelationActionResult
	if err = p.c.Call(ctx, "RelationAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (r *relation.DouyinRelationFollowListResponse, err error) {
	var _args relation.RelationServiceRelationFollowListArgs
	_args.Req = req
	var _result relation.RelationServiceRelationFollowListResult
	if err = p.c.Call(ctx, "RelationFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (r *relation.DouyinRelationFollowerListResponse, err error) {
	var _args relation.RelationServiceRelationFollowerListArgs
	_args.Req = req
	var _result relation.RelationServiceRelationFollowerListResult
	if err = p.c.Call(ctx, "RelationFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
