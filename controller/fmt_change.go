package controller

import (
	"douyin-server/rpc/client"
	"douyin-server/rpc/kitex_gen/feed"
	user "douyin-server/rpc/kitex_gen/user"
	"strconv"
)

func RPCVideo2ControllerVideo(feedVideo *feed.Video) (resp *Video, err error) {
	resp = new(Video)
	if respClient, err := client.UserInfo(feedVideo.AuthorId); err == nil {
		resp = &Video{
			Author:        *RPCUser2ControlUser(respClient),
			ID:            feedVideo.Id,
			PlayURL:       feedVideo.PlayUrl,
			CoverURL:      feedVideo.CoverUrl,
			FavoriteCount: feedVideo.FavoriteCount,
			CommentCount:  feedVideo.CommentCount,
			Title:         feedVideo.Title,
			// TODO is Favorite
		}
	}
	return
}

func RPCUser2ControlUser(userUser *user.User) (resp *User) {
	resp = new(User)
	resp = &User{
		Avatar:          *userUser.Avatar,
		BackgroundImage: *userUser.BackgroundImage,
		FavoriteCount:   *userUser.FavoriteCount,
		FollowCount:     *userUser.FollowCount,
		FollowerCount:   *userUser.FollowerCount,
		ID:              userUser.Id,
		Name:            userUser.Name,
		Signature:       *userUser.Signature,
		TotalFavorited:  strconv.FormatInt(*userUser.TotalFavorited, 10),
		WorkCount:       *userUser.WorkCount,
		// IsFollow:        *userUser.IsFollow,
	}

	return
}
