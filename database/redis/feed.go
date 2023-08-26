package redis

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"encoding/json"
	"fmt"
	"log"
)

func GetVideoById(id int64) (res dao.Dvideo, err error) {
	key := fmt.Sprintf(config.FeedPrefix, id)

	if cnt, err := RedisClient.Exists(Ctx, key).Result(); err == nil && cnt == 1 {
		if tmp, err := RedisClient.Get(Ctx, key).Result(); err == nil {
			if err := json.Unmarshal([]byte(tmp), &res); err != nil {
				log.Println("反序列化失败")
				return res, err
			}
			return res, nil
		}
	}

	//redis缓存中没有该video
	if res, err = dao.GetVideoById(id); err != nil {
		log.Println("Get Video By Id失败")
		return
	}
	//将获取到的结果写入redis
	if err = AddRedisVideoId(key, res); err != nil {
		log.Println("写入缓存失败")
	}
	return res, err
}

func GetVideosByIds(ids []int64) (res []dao.Dvideo, err error) {
	for _, id := range ids {
		tmp, err := GetVideoById(id)
		if err != nil {
			log.Println("Get Video By Ids失败")
			return res, err
		}
		res = append(res, tmp)
	}
	return
}

// 将VideoInfo写入redis中去
func AddRedisVideoId(key string, videoInfo dao.Dvideo) error {
	content, err := json.Marshal(videoInfo)
	if err != nil {
		log.Println("Video序列化失败")
		return err
	}
	if err = RedisClient.Set(Ctx, key, content, config.Exipretime).Err(); err != nil {
		log.Println("video写入redis失败")
		return err
	}
	return nil
}
