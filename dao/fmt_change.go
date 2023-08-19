package dao

import (
	"douyin-server/rpc/kitex_gen/feed"
	"douyin-server/rpc/kitex_gen/user"
)

func DaoVideo2RPCVideo(daoVideo *Dvideo) (resp *feed.Video) {
	resp = new(feed.Video)
	var authorInfo Duser
	authorInfo, _ = GetUserById(daoVideo.Author_id)
	resp = &feed.Video{
		Id:            daoVideo.Id,
		Author:        DaoUser2RPCUser(&authorInfo),
		PlayUrl:       daoVideo.Play_url,
		CoverUrl:      daoVideo.Cover_url,
		UploadTime:    daoVideo.Upload_time.Unix(),
		FavoriteCount: daoVideo.Favorite_count,
		CommentCount:  daoVideo.Comment_count,
		Title:         daoVideo.Title,
		IsFavorite:    daoVideo.Is_favorited,
	}

	return
}

func DaoUser2RPCUser(daoUser *Duser) (resp *user.User) {
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
	return
}
