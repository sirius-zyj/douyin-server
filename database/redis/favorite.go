package redis

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"encoding/json"
	"fmt"
	"log"
)

func GetFavoriteData(userId int64, videoId int64) (res dao.Dfavorite, err error) {
	key := fmt.Sprintf(config.FavoriteDataPrefix, userId, videoId)

	if cnt, err := RedisClient.Exists(Ctx, key).Result(); err == nil && cnt == 1 {
		if tmp, err := RedisClient.Get(Ctx, key).Result(); err == nil {
			if err := json.Unmarshal([]byte(tmp), &res); err != nil {
				log.Println("反序列化失败")
				return res, err
			}
			return res, nil
		}
	}

	//redis缓存中没有该视频的评论信息
	if res, err = dao.GetFavoriteData(userId, videoId); err != nil {
		log.Println("Get FavoriteData失败")
		return
	}
	if res.Id == 0 {
		res.User_id = userId
		res.Video_id = videoId
		res.Action_type = "2"
	}
	//将获取到的结果写入redis
	if err = AddRedisFavoriteData(key, res); err != nil {
		log.Println("写入缓存失败")
	}
	return res, err
}

// 将评论写入redis中去
func AddRedisFavoriteData(key string, favoriteData dao.Dfavorite) error {
	content, err := json.Marshal(favoriteData)
	if err != nil {
		log.Println("FavoriteData序列化失败")
		return err
	}
	if err = RedisClient.Set(Ctx, key, content, config.Exipretime).Err(); err != nil {
		log.Println("FavoriteData写入redis失败")
		return err
	}
	return nil
}
