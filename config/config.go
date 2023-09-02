package config

import (
	"time"
)

// JWT配置数据
var Secret = "douyin"

// mysql配置数据
var (
	Dsn = "root:zyj1314520@tcp(mysql:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
)

// OSS
const (
	OSSAK          = ""
	OSSSK          = ""
	OSSEndPoint    = "oss-cn-hangzhou.aliyuncs.com"
	OSSVideoBucket = "douyin-server-hust"
	OSSImageBucket = "douyin-server-hust-image"
)

// OneDayOfHours 时间
var OneDayOfHours = 60 * 60 * 24
var OneMinute = 60 * 1
var OneMonth = 60 * 60 * 24 * 30
var OneYear = 365 * 60 * 60 * 24

var USE_REDIS = true

// redis的配置数据
var (
	Addr               = "redis:6379"
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

//etcd

const (
	// etcd 端口
	EtcdAddr     = "microservices-etcd:2379"
	FeedAddr     = "0.0.0.0:8880"
	UserAddr     = "0.0.0.0:8881"
	FavoriteAddr = "0.0.0.0:8882"
	CommentAddr  = "0.0.0.0:8883"
	PublishAddr  = "0.0.0.0:8884"
	RelationAddr = "0.0.0.0:8885"
	MessageAddr  = "0.0.0.0:8886"

	// kitex 服务名
	FeedServiceName     = "douyin-server-etcd-feed"
	UserServiceName     = "douyin-server-etcd-user"
	FavoriteServiceName = "douyin-server-etcd-favorite"
	CommentServiceName  = "douyin-server-etcd-comment"
	PublishServiceName  = "douyin-server-etcd-publish"
	RelationServiceName = "douyin-server-etcd-relation"
	MessageServiceName  = "douyin-server-etcd-message"
)
