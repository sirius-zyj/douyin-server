package client

import (
	"context"
	"douyin-server/rpc/kitex_gen/favorite"
	"douyin-server/rpc/kitex_gen/favorite/favoriteservice"
	"log"

	"github.com/cloudwego/kitex/client"
)

var favoriteClient favoriteservice.Client

func initFavoriteClient() {
	c, err := favoriteservice.NewClient("favorite", client.WithHostPorts("0.0.0.0:8882"))
	if err != nil {
		log.Fatal(err)
	}
	favoriteClient = c
}

func ActionFavorite(token string, video_id int64, action_type string) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp = new(favorite.DouyinFavoriteActionResponse)
	resp, err = favoriteClient.FavoriteAction(context.Background(), &favorite.DouyinFavoriteActionRequest{
		Token:      token,
		VideoId:    video_id,
		ActionType: action_type,
	})
	if err != nil {
		log.Printf("ActionFavorite get err %v\n", err)
		return nil, err
	}
	return resp, nil
}

func FavoriteList(user_id int64, token string) (resp *favorite.DouyinFavoriteListResponse, err error) {
	resp = new(favorite.DouyinFavoriteListResponse)
	resp, err = favoriteClient.FavoriteList(context.Background(), &favorite.DouyinFavoriteListRequest{
		UserId: user_id,
		Token:  token,
	})
	if err != nil {
		log.Printf("FavoriteList get err %v\n", err)
		return nil, err
	}
	return resp, nil
}
