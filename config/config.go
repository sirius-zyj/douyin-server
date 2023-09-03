package config

import (
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// JWT配置数据
var Secret string

// etcd
var EtcdAddr string = "microservices-etcd:2379"

// var EtcdAddr string = "microservices-etcd:2379"

var (
	FeedAddr     string
	UserAddr     string
	FavoriteAddr string
	CommentAddr  string
	PublishAddr  string
	RelationAddr string
	MessageAddr  string

	// kitex 服务名
	FeedServiceName     string
	UserServiceName     string
	FavoriteServiceName string
	CommentServiceName  string
	PublishServiceName  string
	RelationServiceName string
	MessageServiceName  string
)

// mysql配置数据
var Dsn string

// redis的配置数据
var (
	RedisAddr          string
	Password           string
	PoolSize           int
	MinConns           int
	FeedPrefix         string        //redis中视频的前缀
	PublishPrefix      string        //redis中视频发布的前缀
	UserPrefix         string        //redis中用户的前缀
	FavoriteDataPrefix string        //redis中收藏信息的前缀
	CommentPrefix      string        //redis中评论的前缀
	FollowDataPrefix   string        //redis中关注信息的前缀
	Exipretime         time.Duration //redis数据的热度消散时间
	OneDayOfHours      = 60 * 60 * 24
	OneMinute          = 60 * 1
	OneMonth           = 60 * 60 * 24 * 30
	OneYear            = 365 * 60 * 60 * 24
	USE_REDIS          = true
)

// OSS
var (
	OSSAK          string
	OSSSK          string
	OSSEndPoint    string
	OSSVideoBucket string
	OSSImageBucket string
)

// OTEL
var (
	RouterOtelName   string
	FeedOtelName     string
	UserOtelName     string
	FavoriteOtelName string
	CommentOtelName  string
	PublishOtelName  string
	RelationOtelName string
	MessageOtelName  string
	JaegerAddr       string
)

func Init() {
	viper.AddRemoteProvider("etcd3", EtcdAddr, "/config/config.yml")
	viper.SetConfigType("yaml")
	if err := viper.ReadRemoteConfig(); err != nil {
		panic(err)
	}

	// jwt
	Secret = viper.GetString("jwt.Secret")
	// etcd
	// EtcdAddr = viper.GetString("etcd.EtcdAddr")
	// microservices
	FeedAddr = viper.GetString("microservices.FeedAddr")
	FeedServiceName = viper.GetString("microservices.FeedServiceName")
	UserAddr = viper.GetString("microservices.UserAddr")
	UserServiceName = viper.GetString("microservices.UserServiceName")
	FavoriteAddr = viper.GetString("microservices.FavoriteAddr")
	FavoriteServiceName = viper.GetString("microservices.FavoriteServiceName")
	CommentAddr = viper.GetString("microservices.CommentAddr")
	CommentServiceName = viper.GetString("microservices.CommentServiceName")
	PublishAddr = viper.GetString("microservices.PublishAddr")
	PublishServiceName = viper.GetString("microservices.PublishServiceName")
	RelationAddr = viper.GetString("microservices.RelationAddr")
	RelationServiceName = viper.GetString("microservices.RelationServiceName")
	MessageAddr = viper.GetString("microservices.MessageAddr")
	MessageServiceName = viper.GetString("microservices.MessageServiceName")
	// redis
	RedisAddr = viper.GetString("redis.RedisAddr")
	Password = viper.GetString("redis.Password")
	PoolSize = viper.GetInt("redis.PoolSize")
	MinConns = viper.GetInt("redis.MinConns")
	FeedPrefix = viper.GetString("redis.FeedPrefix")
	UserPrefix = viper.GetString("redis.UserPrefix")
	FavoriteDataPrefix = viper.GetString("redis.FavoriteDataPrefix")
	CommentPrefix = viper.GetString("redis.CommentPrefix")
	FollowDataPrefix = viper.GetString("redis.FollowDataPrefix")
	Exipretime = viper.GetDuration("redis.Exipretime")
	USE_REDIS = viper.GetBool("redis.USE_REDIS")
	// mysql
	Dsn = viper.GetString("mysql.Dsn")
	// OSS
	OSSAK = viper.GetString("OSS.OSSAK")
	OSSSK = viper.GetString("OSS.OSSSK")
	OSSEndPoint = viper.GetString("OSS.OSSEndPoint")
	OSSVideoBucket = viper.GetString("OSS.OSSVideoBucket")
	OSSImageBucket = viper.GetString("OSS.OSSImageBucket")
	//OTEL
	RouterOtelName = viper.GetString("OTEL.RouterOtelName")
	FeedOtelName = viper.GetString("OTEL.FeedOtelName")
	UserOtelName = viper.GetString("OTEL.UserOtelName")
	FavoriteOtelName = viper.GetString("OTEL.FavoriteOtelName")
	CommentOtelName = viper.GetString("OTEL.CommentOtelName")
	PublishOtelName = viper.GetString("OTEL.PublishOtelName")
	RelationOtelName = viper.GetString("OTEL.RelationOtelName")
	MessageOtelName = viper.GetString("OTEL.MessageOtelName")
	JaegerAddr = viper.GetString("OTEL.JaegerAddr")
}
