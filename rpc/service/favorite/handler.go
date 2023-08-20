package main

import (
	"context"
	"douyin-server/dao"
	favorite "douyin-server/rpc/kitex_gen/favorite"
	"log"
	"strconv"
	"strings"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

func setFavoriteActionResponse(resp *favorite.DouyinFavoriteActionResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp = new(favorite.DouyinFavoriteActionResponse)

	token, video_id, action_type := req.Token, req.VideoId, req.ActionType

	index := strings.Index(token, "*")
	user_id, _ := strconv.ParseInt(token[index+1:], 10, 64)

	video, _ := dao.GetVideoById(video_id)
	fa, err := dao.GetFavoriteData(user_id, video_id)
	if err == nil {
		//获取到的表数据ID为0时代表没有该条点赞数据
		if fa.Id == 0 {
			fa.User_id = user_id
			fa.Action_type = action_type
			fa.Video_id = video_id
			//创建时间
			err = dao.InsertFavorite(fa)
			if err != nil {
				setFavoriteActionResponse(resp, 404, "点赞失败")
			} else {
				dao.UpdateFeed("id", video_id, "favorite_count", 1)         //点赞数+1
				dao.UpdateUser("id", video.Author_id, "total_favorited", 1) //video Author total_favorited+1
				dao.UpdateUser("id", user_id, "favorite_count", 1)          //用户favorite+1
				setFavoriteActionResponse(resp, 0, "点赞成功")
			}
		} else {
			if fa.Action_type != action_type {
				err := dao.ActionFavorite(user_id, video_id, action_type)
				if err != nil {
					setFavoriteActionResponse(resp, 404, "点赞数据更新失败")
				} else {
					if action_type == "1" {
						dao.UpdateFeed("id", video_id, "favorite_count", 1)         //点赞数+1
						dao.UpdateUser("id", video.Author_id, "total_favorited", 1) //video Author total_favorited+1
						dao.UpdateUser("id", user_id, "favorite_count", 1)          //用户favorite+1
					} else {
						dao.UpdateFeed("id", video_id, "favorite_count", -1)         //点赞数-1
						dao.UpdateUser("id", video.Author_id, "total_favorited", -1) //video Author total_favorited+1
						dao.UpdateUser("id", user_id, "favorite_count", -1)          //用户favorite+1
					}
					setFavoriteActionResponse(resp, 0, "点赞数据更新成功")
				}
			} else {
				setFavoriteActionResponse(resp, 0, "Action_type 与数据库中的数据相同")
			}
		}
	} else {
		setFavoriteActionResponse(resp, 404, "点赞发生错误")
	}
	return
}

func setFavoriteListResponse(resp *favorite.DouyinFavoriteListResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	resp = new(favorite.DouyinFavoriteListResponse)
	userID := req.UserId
	VideoIDList, err := dao.GetFavoriteList(userID)
	if err != nil {
		setFavoriteListResponse(resp, 404, "点赞列表查询失败")
	} else {
		var VideoList []dao.Dvideo
		VideoList, err = dao.GetVideosByIds(VideoIDList)
		if err != nil {
			log.Println("查询失败")
			setFavoriteListResponse(resp, 404, "视频数据查询错误")
			return
		}
		setFavoriteListResponse(resp, 0, "视频数据查询成功")
		for _, tmp := range VideoList {
			resp.VideoList = append(resp.VideoList, dao.DaoVideo2RPCVideo(&req.Token, &tmp))
		}
	}
	return
}
