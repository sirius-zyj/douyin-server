package controller

import (
	"douyin-server/rpc/client"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// 获取视频
func Feed(c *gin.Context) {
	var latestTime time.Time
	startTime := c.Query("latest_time")
	if startTime != "0" && startTime != "" {
		s, _ := strconv.ParseInt(startTime, 10, 64)
		latestTime = time.Unix(s, 0)
	} else {
		latestTime = time.Now()
	}
	token := c.PostForm("token")

	if respClient, err := client.Feed(latestTime, token); err == nil {
		var videoList []Video
		for _, tmp := range respClient.VideoList {
			if video, err := RPCVideo2ControllerVideo(tmp); err == nil {
				videoList = append(videoList, *video)
			} else {
				c.JSON(http.StatusExpectationFailed, FeedResponse{})
				return
			}
		}
		log.Println(videoList)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			VideoList: videoList,
			NextTime:  time.Now().Unix(),
		})
	} else {
		c.JSON(http.StatusExpectationFailed, FeedResponse{})
	}
}
