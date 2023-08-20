package dao

import (
	"douyin-server/rpc/kitex_gen/feed"
	"douyin-server/rpc/kitex_gen/user"
	"log"
	"strconv"
	"strings"
)

func DaoVideo2RPCVideo(token *string, daoVideo *Dvideo) (resp *feed.Video) {
	resp = new(feed.Video)
	authorInfo, _ := GetUserById(daoVideo.Author_id)
	resp = &feed.Video{
		Id:            daoVideo.Id,
		Author:        DaoUser2RPCUser(token, &authorInfo),
		PlayUrl:       daoVideo.Play_url,
		CoverUrl:      daoVideo.Cover_url,
		UploadTime:    daoVideo.Upload_time.Unix(),
		FavoriteCount: daoVideo.Favorite_count,
		CommentCount:  daoVideo.Comment_count,
		Title:         daoVideo.Title,
	}
	resp.IsFavorite = CheckWhetherFavorite(token, daoVideo.Id)
	return
}

func DaoUser2RPCUser(token *string, daoUser *Duser) (resp *user.User) {
	resp = new(user.User)
	resp = &user.User{
		Id:              daoUser.ID,
		Avatar:          &daoUser.Avatar,
		Name:            daoUser.Name,
		BackgroundImage: &daoUser.BackgroundImage,
		FavoriteCount:   &daoUser.FavoriteCount,
		FollowCount:     &daoUser.FollowCount,
		FollowerCount:   &daoUser.FollowerCount,
		Signature:       &daoUser.Signature,
		TotalFavorited:  &daoUser.TotalFavorited,
		WorkCount:       &daoUser.WorkCount,
	}
	resp.IsFollow = CheckWhetherFollow(token, daoUser.ID)
	log.Println("resp.IsFollow: ", resp.IsFollow)
	return
}

func CheckWhetherFavorite(token *string, video_id int64) bool {
	if token != nil {
		index := strings.Index(*token, "*")
		user_id, _ := strconv.ParseInt((*token)[index+1:], 10, 64)
		// 判断是否已经点赞
		favoriteData, err := GetFavoriteData(user_id, video_id)
		if err == nil && favoriteData.Id != 0 && favoriteData.Action_type == "1" {
			return true
		}
	}
	return false
}

func CheckWhetherFollow(token *string, follow_id int64) bool {
	if token != nil {
		index := strings.Index(*token, "*")
		user_id, _ := strconv.ParseInt((*token)[index+1:], 10, 64)
		// 判断是否已经关注
		followData, err := GetFollowData(user_id, follow_id)
		if err == nil && followData.Id != 0 && followData.Action_type == "1" {
			return true
		}
	}
	return false
}
