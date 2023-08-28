package database

import (
	"douyin-server/database/dao"
	"douyin-server/middleware/jwt"
	"douyin-server/rpc/kitex_gen/feed"
	"douyin-server/rpc/kitex_gen/user"
)

func DaoVideo2RPCVideo(token *string, daoVideo *dao.Dvideo) (resp *feed.Video) {
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

func DaoUser2RPCUser(token *string, daoUser *dao.Duser) (resp *user.User) {
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
	return
}

func CheckWhetherFavorite(token *string, videoId int64) bool {
	if token != nil {
		userId := jwt.GetUserIdByToken(*token)
		// 判断是否已经点赞
		favoriteData, err := GetFavoriteData(userId, videoId)
		if err == nil && favoriteData.Id != 0 && favoriteData.Action_type == "1" {
			return true
		}
	}
	return false
}

func CheckWhetherFollow(token *string, followId int64) bool {
	if token != nil {
		userId := jwt.GetUserIdByToken(*token)
		// 判断是否已经关注
		followData, err := GetFollowData(userId, followId)
		if err == nil && followData.Id != 0 && followData.Action_type == "1" {
			return true
		}
	}
	return false
}
