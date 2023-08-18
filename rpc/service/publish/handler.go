package main

import (
	"context"
	"douyin-server/dao"
	"douyin-server/rpc/kitex_gen/feed"
	publish "douyin-server/rpc/kitex_gen/publish"
	"log"
	"strconv"
	"strings"
	"time"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// Publish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) Publish(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	resp = new(publish.DouyinPublishActionResponse)

	video_Data := []byte(req.Data)
	token, video_title := req.Token, req.Title
	playUrl, coverUrl, err := dao.UploadVideo(&video_Data)
	if err != nil {
		log.Println("上传视频失败")
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "视频上传失败"
		return
	}
	//------创建视频--------
	index := strings.Index(token, "*")
	user_id, _ := strconv.ParseInt(token[index+1:], 10, 64)
	video := dao.Dvideo{
		Author_id:   user_id,
		Play_url:    playUrl,
		Cover_url:   coverUrl,
		Upload_time: time.Now(),
		Title:       video_title,
	}
	if Insert_error := dao.InsertVideo(video); Insert_error != nil {
		log.Println("Insert_error: ", Insert_error)
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = err.Error()
		return
	}
	//---------------------------
	resp.StatusCode = 0
	resp.StatusMsg = new(string)
	*resp.StatusMsg = " uploaded successfully"
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	resp = new(publish.DouyinPublishListResponse)

	userID := req.UserId
	//获取目标用户的所有作品将其传递给APP
	if videoList, err := dao.GetVideoByUserId(userID); err != nil {
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = err.Error()
		return resp, err
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "success"
		for _, tmp := range videoList {
			resp.VideoList = append(resp.VideoList, &feed.Video{
				Id:       tmp.Id,
				PlayUrl:  tmp.Play_url,
				CoverUrl: &tmp.Cover_url,
				Title:    &tmp.Title,
			})
		}
	}
	return
}
