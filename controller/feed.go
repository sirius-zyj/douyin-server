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
  startTime := c.Query("latest_time")
  log.Printf("请求的时间" + startTime)

  
  var lastTime time.Time
  if startTime != "0" {
    s , _ := strconv.ParseInt(startTime , 10 , 64)
    lastTime = time.Unix(s , 0)
  }else{
    lastTime = time.Now()
  }
  log.Printf("请求的时间戳 %v" , lastTime)

  VideoQueue := GetVideo(lastTime)
  // VideoQueue := DemoVideos
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: VideoQueue,
		NextTime:  time.Now().Unix(),
	})
}

// 获取查询到的视频切片
func GetVideo(userId int64, lastTime time.Time) []Video {
	resp, err := client.GetVideoByUserId(userId)
	// resp, err := dao.GetVideoByTime(lastTime)
	var ans VideoSlice
	if err != nil {
		return ans
	}
	for _, temp := range resp {
		var video Video
		video.Dvideo = temp
		ans.Append(video)
	}
	return ans
}

func (vs *VideoSlice) Append(video Video) {
	*vs = append(*vs, video)
}
