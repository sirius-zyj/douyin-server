package redis

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"encoding/json"
	"fmt"
	"log"
)

// 查询关注信息
func GetFollowData(userId int64, followId int64) (res dao.Dfollow, err error) {
	key := fmt.Sprintf(config.FollowDataPrefix, userId, followId)

	if cnt, err := RedisClient.Exists(Ctx, key).Result(); err == nil && cnt == 1 {
		if tmp, err := RedisClient.Get(Ctx, key).Result(); err == nil {
			if err = json.Unmarshal([]byte(tmp), &res); err != nil {
				log.Println("反序列化失败")
				return res, err
			}
			return res, nil
		}
	}

	//redis缓存中没有该视频的评论信息
	if res, err = dao.GetFollowData(userId, followId); err != nil {
		log.Println("Get FollowData失败")
		return
	}
	//将获取到的结果写入redis
	if res.Id != 0 {
		if err = AddRedisFollowData(&key, res); err != nil {
			log.Println("写入缓存失败")
		}
	}
	return res, err
}

// 将评论写入redis中去
func AddRedisFollowData(key *string, followData dao.Dfollow) error {
	content, err := json.Marshal(followData)
	if err != nil {
		log.Println("FollowData序列化失败")
		return err
	}
	if err = RedisClient.Set(Ctx, *key, content, config.Exipretime).Err(); err != nil {
		log.Println("FollowData写入redis失败")
		return err
	}
	return nil
}
