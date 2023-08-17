package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"douyin-server/dao"

	"github.com/gin-gonic/gin"
)

// FavoriteAction 点赞或者取消点赞
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	index := strings.Index(token, "*")
	useid := token[index+1:]
	// log.Println(useid)
	user_id, _ := strconv.ParseInt(useid, 10, 64)
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	var action bool
	if action_type == 1 {
		action = true
	} else {
		action = false
	}
	fa, err := dao.GetFavoriteData(user_id, video_id)
	if err == nil {
		//获取到的表数据ID为0时代表没有该条点赞数据
		log.Println(fa.Id)
		if fa.Id == 0 {
			fa.User_id = user_id
			fa.Cancel = action
			fa.Video_id = video_id
			//创建时间
			err := dao.InsertFavorite(fa)
			if err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "点赞失败"})
			}
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else {
			err := dao.ActionFavorite(user_id, video_id, action)
			if err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "点赞数据更新失败"})
			} else {
				c.JSON(http.StatusOK, Response{StatusCode: 0})
			}
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "点赞发生错误"})
	}
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	userid := c.Query("user_id")
	userID, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		log.Println("数据转换错误")
	}
	VideoIDList, err := dao.GetFavoriteList(userID)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "视频ID查询错误",
			},
		})
	} else {
		VideoList, err := dao.GetVideosByIds(VideoIDList)
		if err != nil {
			log.Println("查询失败")
			c.JSON(http.StatusOK, VideoListResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "视频数据查询错误",
				},
			})
			return
		}
		var videolist VideoSlice
		for _, tmp := range VideoList {
			var v Video
			v.Dvideo = tmp
			//------还有获取点赞数，获取评论数
			videolist.Append(v)
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: videolist,
		})
	}
}
