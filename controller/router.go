package controller

import (
	"douyin-server/middleware/validate"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//sensitive filter
	validate.InitSensitive()

	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	apiRouter.Use(validate.InitRateLimit(time.Second, 888, 888))

	// basic apis
	apiRouter.GET("/feed/", Feed)
	apiRouter.GET("/user/", UserInfo)
	apiRouter.POST("/user/register/", validate.RegisterWordsFilter, Register)
	apiRouter.POST("/user/login/", Login)
	apiRouter.POST("/publish/action/", validate.InitRateLimit(time.Minute, 10, 1), validate.PublishActionWordsFilter, validate.VideoFilter, Publish)
	apiRouter.GET("/publish/list/", PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", FavoriteAction)
	apiRouter.GET("/favorite/list/", FavoriteList)
	apiRouter.POST("/comment/action/", validate.CommentActionWordsFilter, CommentAction)
	apiRouter.GET("/comment/list/", CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", RelationAction)
	apiRouter.GET("/relation/follow/list/", FollowList)
	apiRouter.GET("/relation/follower/list/", FollowerList)
	apiRouter.GET("/relation/friend/list/", FriendList)
	apiRouter.GET("/message/chat/", MessageChat)
	apiRouter.POST("/message/action/", validate.MessageActionWordsFilter, MessageAction)
}
