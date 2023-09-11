package main

import (
	// 其他导入语句

	"context"
	"douyin-server/config"
	"douyin-server/database/dao"
	"douyin-server/middleware/gorse"
	"log"
)

func main1() {
	config.MasterDsn = "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	config.SlaveDsn = "root:123456@tcp(127.0.0.1:3307)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	dao.Init()
	// videos, _ := dao.GetVideoByAuthorId(873378760114176)
	// for _, v := range videos {
	// 	if err := gorse.GetInstance().PublishToGorse(context.Background(), &v); err != nil {
	// 		log.Println(err)
	// 	}
	// }
	video, _ := dao.GetVideoById(2283494503890944)
	if err := gorse.PublishToGorse(context.Background(), &video); err != nil {
		log.Println(err)
	}
	// user, _ := dao.GetUserById(1305490470486016)
	// gorse.GetInstance().RegisterToGorse(context.Background(), &user)
	// fav, _ := dao.GetFavoriteList(869002494623744)
	// for _, i := range fav {
	// 	fa, _ := dao.GetFavoriteData(869002494623744, i)
	// 	gorse.GetInstance().FavoriteToGorse(context.Background(), &fa)
	// }
	// gorse.GetInstance().FavoriteToGorse(context.Background(), &fa)
	// gorse.GetUserRecommend(context.Background(), 869002494623744)
	// if _, err := gorse.GetInstance().InsertFeedback(context.Background(), []gorse.Feedback{
	// 	{FeedbackType: "star", UserId: "bob", ItemId: "vuejs:vue", Timestamp: "2022-02-24"},
	// }); err != nil {
	// 	log.Println("Insert Err", err)
	// }

	// if resp, err := gorse.GetInstance().GetRecommend(context.Background(), "bob", "", 10); err != nil {
	// 	log.Println("Get Err", err)
	// } else {
	// 	log.Println(resp)
	// }
}
