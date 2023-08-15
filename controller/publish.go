package controller

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish 视频投稿
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
  //文件存储到的位置
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

//查看他人的作品列表
func PublishList(c *gin.Context) {
  user_ID , _ := c.GetQuery("user_id")
  userID , _ := strconv.ParseInt(user_ID , 10 , 64)
  log.Printf("获取到的目标用户id %v" , userID)

  //获取目标用户的所有作品将其传递给APP
  if _ , err := dao.GetVideoByUserId(userID); err != nil {
    c.JSON(http.StatusOK, VideoListResponse{
		    Response: Response{
			  StatusCode: 1,
        StatusMsg: "查询出错",
		  },
	  })
  }else {
    
    c.JSON(http.StatusOK, VideoListResponse{
    	Response: Response{
    			StatusCode: 0,
    	},
  		VideoList: DemoVideos,
	  })
  }
}
