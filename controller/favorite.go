package controller

import (
	"log"
	"net/http"
	"strconv"

	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

type FavoriteActionRequest struct {
	ActionType string `form:"action_type"` // 1-点赞，2-取消点赞
	Token      string `form:"token"`       // 用户鉴权token
	VideoID    string `form:"video_id"`    // 视频id
}

type FavoriteListRequest struct {
	Token  string `form:"token"`   // 用户鉴权token
	UserID string `form:"user_id"` // 用户id
}

// FavoriteAction 点赞或者取消点赞
func FavoriteAction(c *gin.Context) {
	var req FavoriteActionRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("FavoriteActionRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserResponse{Response: Response{StatusCode: 404}})
		return
	}
	video_id, _ := strconv.ParseInt(req.VideoID, 10, 64)

	if respClient, err := client.ActionFavorite(req.Token, video_id, req.ActionType); err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)})
	} else {
		c.JSON(http.StatusInternalServerError, Response{})
	}
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	var req FavoriteListRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("FavoriteListRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserResponse{Response: Response{StatusCode: 404}})
		return
	}
	userID, _ := strconv.ParseInt(req.UserID, 10, 64)

	if respClient, err := client.FavoriteList(userID); err == nil {
		var videoList []Video
		for _, tmp := range respClient.VideoList {
			if video, err := RPCVideo2ControllerVideo(tmp); err == nil {
				videoList = append(videoList, *video)
			} else {
				c.JSON(http.StatusServiceUnavailable, FeedResponse{Response: Response{StatusCode: 404, StatusMsg: "RPC Video2ControllerVideo错误"}})
				return
			}
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response:  Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			VideoList: videoList,
		})

	} else {
		c.JSON(http.StatusInternalServerError, VideoListResponse{})
	}
}
