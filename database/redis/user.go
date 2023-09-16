package redis

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"encoding/json"
	"fmt"
	"log"
)

func GetUserById(userId int64) (res dao.Duser, err error) {
	key := fmt.Sprintf(config.UserPrefix, userId)

	if cnt, err := RedisClient.Exists(Ctx, key).Result(); err == nil && cnt == 1 {
		if tmp, err := RedisClient.Get(Ctx, key).Result(); err == nil {
			if err := json.Unmarshal([]byte(tmp), &res); err != nil {
				log.Println("反序列化失败")
				return res, err
			}
			RedisClient.Expire(Ctx, key, GenExpireTime())
			return res, nil
		}
	}

	//redis缓存中没有该user
	if res, err = dao.GetUserById(userId); err != nil {
		log.Println("Get User By Id失败")
		return
	}
	//将获取到的结果写入redis
	if err = AddRedisUserId(key, res); err != nil {
		log.Println("写入缓存失败")
	}
	return res, err
}

// 将UserInfo写入redis中去
func AddRedisUserId(key string, userInfo dao.Duser) error {
	content, err := json.Marshal(userInfo)
	if err != nil {
		log.Println("User序列化失败")
		return err
	}
	if err = RedisClient.Set(Ctx, key, content, GenExpireTime()).Err(); err != nil {
		log.Println("user写入redis失败")
		return err
	}
	return nil
}

func EraseRedisUserId(key string) (err error) {
	RedisClient.Del(Ctx, key)
	return nil
}
