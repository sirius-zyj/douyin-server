package main

import (
	"context"
	"douyin-server/database"
	"douyin-server/database/dao"
	"douyin-server/middleware/jwt"
	publish "douyin-server/rpc/kitex_gen/publish"
	"log"
	"time"
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

	token, video_title := req.Token, req.Title
	playUrl, coverUrl, err := dao.UploadVideo(&req.Data)
	if err != nil {
		log.Println("上传视频失败")
		setPublishActionResp(resp, 404, "上传视频失败")
		return
	}
	//------创建视频--------
	userId := jwt.GetUserIdByToken(token)

	video := dao.Dvideo{
		Id:          dao.SnowFlakeNode.Generate().Int64(),
		Author_id:   userId,
		Play_url:    playUrl,
		Cover_url:   coverUrl,
		Upload_time: time.Now(),
		Title:       video_title,
	}
	if err = dao.Tran_InsertVideo(video); err != nil {
		log.Println("Insert_error: ", err)
		setPublishActionResp(resp, 404, err.Error())
		return
	}
	//---------------------------
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
