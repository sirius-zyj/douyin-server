package controller

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"douyin-server/dao"
	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish 视频投稿
func Publish(c *gin.Context) {
	token, video_title := c.PostForm("token"), c.Query("title")

	file, _ := c.FormFile("data")
	// 打开上传的文件
	src, _ := file.Open()
	defer src.Close()
	// 读取文件内容
	video_Data, _ := ioutil.ReadAll(src)

	if respClient, err := client.Publish(token, video_Data, video_title); err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)})
	} else {
		c.JSON(http.StatusExpectationFailed, Response{})
	}
}

// 根据用户ID查找该用户作品列表
func PublishList(c *gin.Context) {
	user_ID, _ := c.GetQuery("user_id")
	userID, _ := strconv.ParseInt(user_ID, 10, 64)
	token := c.PostForm("token")

	if respClient, err := client.PublishList(userID, token); err == nil {
		log.Println("PublishList: ", *respClient)
		var videoList VideoSlice
		for _, video := range respClient.VideoList {
			V := dao.Dvideo{
				Id:        video.Id,
				Play_url:  video.PlayUrl,
				Cover_url: *video.CoverUrl,
				Title:     *video.Title,
			}
			videoList.Append(Video{V, 0, 0, false})
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response:  Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			VideoList: videoList,
		})
	} else {
		c.JSON(http.StatusExpectationFailed, VideoListResponse{})
	}

}
