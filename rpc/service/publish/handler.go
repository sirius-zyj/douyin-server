package main

import (
	"context"
	"douyin-server/database"
	"douyin-server/database/dao"
	"douyin-server/middleware/rabbitmq"
	publish "douyin-server/rpc/kitex_gen/publish"
	"io/ioutil"
	"os"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

func setPublishActionResp(resp *publish.DouyinPublishActionResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// Publish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) Publish(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	resp = new(publish.DouyinPublishActionResponse)

	fileName := dao.SnowFlakeNode.Generate().String()
	if err = ioutil.WriteFile(fileName, req.Data, 0644); err != nil {
		setPublishActionResp(resp, 404, "上传失败")
		return
	}

	if err = rabbitmq.AddToMQ(rabbitmq.PublishActionMessage{
		Filename: fileName,
		Token:    req.Token,
		Title:    req.Title,
	}); err != nil {
		os.Remove(fileName)
		setPublishActionResp(resp, 404, "上传失败")
		return
	}

	setPublishActionResp(resp, 0, "上传成功")
	return
}

func setPublishListResp(resp *publish.DouyinPublishListResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	resp = new(publish.DouyinPublishListResponse)

	//获取目标用户的所有作品将其传递给APP
	if videoList, err := database.GetVideoByAuthorId(req.UserId); err != nil {
		setPublishListResp(resp, 404, err.Error())
		return resp, err
	} else {
		setPublishListResp(resp, 0, "success")
		for _, tmp := range videoList {
			resp.VideoList = append(resp.VideoList, database.DaoVideo2RPCVideo(&req.Token, &tmp))
		}
	}
	return
}
