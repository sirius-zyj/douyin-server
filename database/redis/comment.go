package redis

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// func PublishComment()

// func DeleteComment()

func GetAllComments(videoId int64) (commentList []dao.Dcomments, err error) {
	key := fmt.Sprintf(config.CommentPrefix, videoId)

	if cnt, err := RedisClient.Exists(Ctx, key).Result(); err == nil && cnt == 1 {
		if res, err := RedisClient.HGetAll(Ctx, key).Result(); err == nil {
			for _, tmp := range res {
				var comment dao.Dcomments
				err := json.Unmarshal([]byte(tmp), &comment)
				if err != nil {
					log.Println("反序列化失败")
					return commentList, err
				}
				commentList = append(commentList, comment)
			}
			return commentList, nil
		}
	}
	//redis缓存中没有该视频的评论信息
	commentList, err = dao.GetAllComments(videoId)
	if err != nil {
		log.Println("拉取评论失败")
		return commentList, err
	}
	//将获取到的结果写入redis
	err1 := AddRedisComment(key, commentList)
	if err1 != nil {
		log.Println("写入缓存失败")
	}
	return commentList, nil
}

// 将评论写入redis中去
func AddRedisComment(key string, commentList []dao.Dcomments) error {
	commentMap := make(map[string]interface{})
	for _, tmp := range commentList {
		id := strconv.Itoa(int(tmp.Id))
		content, err := json.Marshal(tmp)
		if err != nil {
			log.Println("评论序列化失败")
			panic(err)
		}
		commentMap[id] = content
	}
	if len(commentMap) == 0 {
		return nil
	}
	err := RedisClient.HMSet(Ctx, key, commentMap).Err()
	if err != nil {
		log.Println("评论写入redis失败")
		return nil
	}
	RedisClient.Expire(Ctx, key, config.Exipretime)
	return nil
}
