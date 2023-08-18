package main

import (
	"context"
	"douyin-server/dao"
	favorite "douyin-server/rpc/kitex_gen/favorite"
	"douyin-server/rpc/kitex_gen/feed"
	"log"
	"strconv"
	"strings"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp = new(favorite.DouyinFavoriteActionResponse)

	token, video_id, action_type := req.Token, req.VideoId, req.ActionType

	index := strings.Index(token, "*")
	user_id, _ := strconv.ParseInt(token[index+1:], 10, 64)

	var action bool
	if action_type == 1 {
		action = true
	} else {
		action = false
	}
	fa, err := dao.GetFavoriteData(user_id, video_id)
	if err == nil {
		//获取到的表数据ID为0时代表没有该条点赞数据
		log.Println(fa.Id)
		if fa.Id == 0 {
			fa.User_id = user_id
			fa.Cancel = action
			fa.Video_id = video_id
			//创建时间
			err = dao.InsertFavorite(fa)
			if err != nil {
				resp.StatusCode = 404
				resp.StatusMsg = new(string)
				*resp.StatusMsg = "点赞失败"
			} else {
				resp.StatusCode = 0
				resp.StatusMsg = new(string)
				*resp.StatusMsg = "点赞成功"
			}
		} else {
			err := dao.ActionFavorite(user_id, video_id, action)
			if err != nil {
				resp.StatusCode = 404
				resp.StatusMsg = new(string)
				*resp.StatusMsg = "点赞数据更新失败"
			} else {
				resp.StatusCode = 0
				resp.StatusMsg = new(string)
				*resp.StatusMsg = "点赞update成功"
			}
		}
	} else {
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "点赞发生错误"
	}
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	resp = new(favorite.DouyinFavoriteListResponse)
	userID := req.UserId
	VideoIDList, err := dao.GetFavoriteList(userID)
	if err != nil {
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "视频ID查询错误"
	} else {
		var VideoList []dao.Dvideo
		VideoList, err = dao.GetVideosByIds(VideoIDList)
		if err != nil {
			log.Println("查询失败")
			resp.StatusCode = 404
			resp.StatusMsg = new(string)
			*resp.StatusMsg = "视频数据查询错误"
			return
		}
		resp.StatusCode = 0
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "视频数据查询成功"
		for _, tmp := range VideoList {
			v := feed.Video{
				Id:      tmp.Id,
				PlayUrl: tmp.Play_url,
			}
			resp.VideoList = append(resp.VideoList, &v)
		}
	}
	return
}
