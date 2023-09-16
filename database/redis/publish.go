package redis

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func GetVideoByAuthorId(id int64) (videoList []dao.Dvideo, err error) {
	key := fmt.Sprintf(config.PublishPrefix, id)
	if cnt, err := RedisClient.Exists(Ctx, key).Result(); err == nil && cnt == 1 {
		if res, err := RedisClient.HGetAll(Ctx, key).Result(); err == nil {
			for _, tmp := range res {
				var video dao.Dvideo
				err := json.Unmarshal([]byte(tmp), &video)
				if err != nil {
					log.Println("反序列化失败")
					return videoList, err
				}
				videoList = append(videoList, video)
			}
			RedisClient.Expire(Ctx, key, GenExpireTime())
			return videoList, nil
		}
	}
	//redis缓存中没有该视频的评论信息
	videoList, err = dao.GetVideoByAuthorId(id)
	if err != nil {
		log.Println("拉取video失败")
		return videoList, err
	}
	//将获取到的结果写入redis
	err1 := AddRedisPublish(key, videoList)
	if err1 != nil {
		log.Println("写入缓存失败")
	}
	return videoList, nil
}

// 将评论写入redis中去
func AddRedisPublish(key string, videoList []dao.Dvideo) error {
	videoMap := make(map[string]interface{})
	for _, tmp := range videoList {
		id := strconv.Itoa(int(tmp.Id))
		content, err := json.Marshal(tmp)
		if err != nil {
			log.Println("video序列化失败")
			panic(err)
		}
		videoMap[id] = content
	}
	if len(videoMap) == 0 {
		return nil
	}
	err := RedisClient.HMSet(Ctx, key, videoMap).Err()
	if err != nil {
		log.Println("评论写入redis失败")
		return nil
	}
	RedisClient.Expire(Ctx, key, GenExpireTime())
	return nil
}

func EraseRedisPublish(key string) (err error) {
	RedisClient.Del(Ctx, key)
	return nil
}
