package database

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"douyin-server/database/redis"
	"fmt"
)

func Init() {
	dao.Init()
	redis.InitRedis()
}

func GetVideoById(videoId int64) (dao.Dvideo, error) {
	if config.USE_REDIS {
		return redis.GetVideoById(videoId)
	}
	return dao.GetVideoById(videoId)
}

func GetUserById(userId int64) (dao.Duser, error) {
	if config.USE_REDIS {
		return redis.GetUserById(userId)
	}
	return dao.GetUserById(userId)
}

func GetVideosByIds(ids []int64) ([]dao.Dvideo, error) {
	if config.USE_REDIS {
		return redis.GetVideosByIds(ids)
	}
	return dao.GetVideosByIds(ids)
}

func GetVideoByAuthorId(id int64) ([]dao.Dvideo, error) {
	if config.USE_REDIS {
		return redis.GetVideoByAuthorId(id)
	}
	return dao.GetVideoByAuthorId(id)
}

func GetFavoriteData(userId int64, videoId int64) (dao.Dfavorite, error) {
	if config.USE_REDIS {
		return redis.GetFavoriteData(userId, videoId)
	}
	return dao.GetFavoriteData(userId, videoId)
}

func InsertFavorite(faDate *dao.Dfavorite, authorId int64) (err error) {
	key := fmt.Sprintf(config.FavoriteDataPrefix, faDate.User_id, faDate.Video_id) //
	err = dao.Tran_InsertFavorite(faDate, authorId)
	if config.USE_REDIS && err == nil {
		redis.AddRedisFavoriteData(key, *faDate)
	}
	return
}

func GetAllComments(videoId int64) ([]dao.Dcomments, error) {
	if config.USE_REDIS {
		return redis.GetAllComments(videoId)
	}
	return dao.GetAllComments(videoId)
}

func GetFollowData(userId int64, followId int64) (res dao.Dfollow, err error) {
	if config.USE_REDIS {
		return redis.GetFollowData(userId, followId)
	}
	return dao.GetFollowData(userId, followId)
}

func InsertFollow(foDate *dao.Dfollow) (err error) {
	key := fmt.Sprintf(config.FollowDataPrefix, foDate.User_id, foDate.Follow_id)
	err = dao.Tran_InsertFollow(foDate)
	if config.USE_REDIS && err == nil {
		redis.AddRedisFollowData(&key, *foDate)
	}
	return
}
