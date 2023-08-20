package controller

import (
	"douyin-server/rpc/client"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RelationActionRequest struct {
	ActionType string `form:"action_type"` // 1-关注，2-取消关注
	FollowId   string `form:"to_user_id"`  // 对方用户id
	Token      string `form:"token"`       // 用户鉴权token
}

type RelationListRequest struct {
	Token  string `json:"token"`   // 用户鉴权token
	UserID string `json:"user_id"` // 用户id
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction 处理关注和取消关注
func RelationAction(c *gin.Context) {
	var req RelationActionRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("RelationActionRequest Err : ", err)
		c.JSON(http.StatusBadRequest, FeedResponse{Response: Response{StatusCode: 404, StatusMsg: "参数错误"}})
		return
	}

	follow_id, _ := strconv.ParseInt(req.FollowId, 10, 64)
	if respClient, err := client.RelationAction(req.Token, follow_id, req.ActionType); err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)})
	} else {
		c.JSON(http.StatusInternalServerError, Response{})
	}

}

// FollowList 获取关注列表
func FollowList(c *gin.Context) {
	// var req RelationListRequest
	// if err := c.ShouldBind(&req); err != nil {
	// 	log.Println("RelationListRequest Err : ", err)
	// 	c.JSON(http.StatusBadRequest, FeedResponse{Response: Response{StatusCode: 404, StatusMsg: "参数错误"}})
	// 	return
	// }

	// user_id, _ := strconv.ParseInt(req.UserID, 10, 64)

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FollowerList 获取粉丝列表
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FriendList 获取朋友列表
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
