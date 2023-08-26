package config

import (
	"time"
)

// JWT配置数据
var Secret = "douyin"

// mysql配置数据
var (
	Dsn = "root:SgIopgnq@tcp(172.16.32.35:51007)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
)

// OneDayOfHours 时间
var OneDayOfHours = 60 * 60 * 24
var OneMinute = 60 * 1
var OneMonth = 60 * 60 * 24 * 30
var OneYear = 365 * 60 * 60 * 24

var USE_REDIS = true

// redis的配置数据
var (
	Addr               = "127.0.0.1:6379"
	Password           = ""
	PoolSize           = 30
	MinConns           = 30
	FeedPrefix         = "feed{id:%d}"                       //redis中视频的前缀
	PublishPrefix      = "publish{author_id:%d}"             //redis中视频发布的前缀
	UserPrefix         = "user{id:%d}"                       //redis中用户的前缀
	FavoriteDataPrefix = "favorite{user_id:%d; video_id:%d}" //redis中收藏信息的前缀
	CommentPrefix      = "comment{video_id:%d}"              //redis中评论的前缀
	FollowDataPrefix   = "follow{user_id:%d; follow_id:%d}"  //redis中关注信息的前缀
	Exipretime         = time.Second * 120                   //redis数据的热度消散时间
)
