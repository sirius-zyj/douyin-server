package controller

import (
	"log"
	"net/http"
	"strconv"

	"douyin-server/dao"
	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

// FavoriteAction 点赞或者取消点赞
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)

	if respClient, err := client.ActionFavorite(token, video_id, int32(action_type)); err == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: respClient.StatusCode,
			StatusMsg: func() string {
				if respClient.StatusMsg != nil {
					return *respClient.StatusMsg
				}
				return ""
			}(),
		})
	} else {
		c.JSON(http.StatusExpectationFailed, Response{})
	}
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	userid := c.Query("user_id")
	userID, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		log.Println("数据转换错误")
	}

	if respClient, err := client.FavoriteList(userID); err == nil {
		log.Println(*respClient.StatusMsg)
		var videolist VideoSlice
		for _, tmp := range respClient.VideoList {
			var v Video
			v.Dvideo = dao.Dvideo{
				Id:       tmp.Id,
				Play_url: tmp.PlayUrl,
			}
			//------还有获取点赞数，获取评论数
			videolist.Append(v)
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: respClient.StatusCode,
				StatusMsg:  StatusMsg(respClient.StatusMsg),
			},
			VideoList: videolist,
		})

	} else {
		c.JSON(http.StatusExpectationFailed, VideoListResponse{})
	}
}
