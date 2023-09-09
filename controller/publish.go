package controller

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

type PublishListRequest struct {
	Token  string `form:"token"`   // 用户鉴权token
	UserID string `form:"user_id"` // 用户id
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish 视频投稿
func Publish(c *gin.Context) {
	token, video_title := c.PostForm("token"), c.PostForm("title")

	file, err := c.FormFile("data")
	if err != nil {
		log.Println("FormFile Err : ", err)
		c.JSON(http.StatusBadRequest, Response{StatusCode: 404, StatusMsg: "参数错误"})
		return
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		log.Println("OpenFile Err : ", err)
		c.JSON(http.StatusBadRequest, Response{StatusCode: 404, StatusMsg: "file open error"})
		return
	}
	defer src.Close()
	// 读取文件内容
	video_Data, err := ioutil.ReadAll(src)
	if err != nil {
		log.Println("ReadAll Err : ", err)
		c.JSON(http.StatusBadRequest, Response{StatusCode: 404, StatusMsg: "file read error"})
		return
	}

	if respClient, err := client.Publish(token, video_Data, video_title); err == nil {
		log.Println("Publish StatusCode : ", respClient.StatusCode, StatusMsg(respClient.StatusMsg))
		c.JSON(http.StatusOK, Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)})
	} else {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, Response{})
	}
}

// 根据用户ID查找该用户作品列表
func PublishList(c *gin.Context) {
	var req PublishListRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("PublishListRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserResponse{Response: Response{StatusCode: 404}})
		return
	}
	if req.Token == "" {
		c.JSON(http.StatusOK, VideoListResponse{Response: Response{StatusCode: 200}})
		return
	}
	userID, _ := strconv.ParseInt(req.UserID, 10, 64)

	if respClient, err := client.PublishList(userID, req.Token); err == nil {
		var videoList []Video

		for _, tmp := range respClient.VideoList {
			//------还有获取点赞数，获取评论数
			if video, err := RPCVideo2ControllerVideo(tmp); err == nil {
				videoList = append(videoList, *video)
			} else {
				c.JSON(http.StatusInternalServerError, VideoListResponse{})
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
