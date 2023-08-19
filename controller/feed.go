package controller

import (
	"douyin-server/rpc/client"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedRequest struct {
	LatestTime *string `form:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `form:"token,omitempty"`       // 用户登录状态下设置
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// 获取视频
func Feed(c *gin.Context) {
	var req FeedRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("FeedRequest Err : ", err)
		c.JSON(http.StatusBadRequest, FeedResponse{Response: Response{StatusCode: 404, StatusMsg: "参数错误"}})
		return
	}

	var latestTime time.Time
	if req.LatestTime != nil && *req.LatestTime != "0" && *req.LatestTime != "" {
		s, _ := strconv.ParseInt(*req.LatestTime, 10, 64)
		latestTime = time.Unix(s, 0)
	} else {
		latestTime = time.Now()
	}
	var token string
	if req.Token != nil {
		token = *req.Token
	}

	if respClient, err := client.Feed(latestTime, token); err == nil {
		var videoList []Video
		for _, tmp := range respClient.VideoList {
			if video, err := RPCVideo2ControllerVideo(tmp); err == nil {
				videoList = append(videoList, *video)
			} else {
				c.JSON(http.StatusServiceUnavailable, FeedResponse{Response: Response{StatusCode: 404, StatusMsg: "RPC Video2ControllerVideo错误"}})
				return
			}
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			VideoList: videoList,
			NextTime:  time.Now().Unix(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, FeedResponse{Response: Response{StatusCode: 404, StatusMsg: "RPC错误"}}) //500 RPC错误
	}
}
