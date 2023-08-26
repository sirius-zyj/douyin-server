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

type UserListRequest struct {
	Token  string `form:"token"`   // 用户鉴权token
	UserID string `form:"user_id"` // 用户id
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
		c.JSON(http.StatusBadRequest, UserListResponse{Response: Response{StatusCode: 404, StatusMsg: "参数错误"}})
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
	var req UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("FollowListRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserListResponse{Response: Response{StatusCode: 404, StatusMsg: "参数错误"}})
		return
	}

	user_id, _ := strconv.ParseInt(req.UserID, 10, 64)

	if respClient, err := client.RelationFollowList(req.Token, user_id); err == nil {
		var userList []User
		for _, tmp := range respClient.UserList {
			userList = append(userList, *RPCUser2ControllerUser(tmp))
		}
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			UserList: userList,
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{})
	}
}

// FollowerList 获取粉丝列表
func FollowerList(c *gin.Context) {
	var req UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("FollowerListRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserListResponse{Response: Response{StatusCode: 404, StatusMsg: "参数错误"}})
		return
	}

	user_id, _ := strconv.ParseInt(req.UserID, 10, 64)

	if respClient, err := client.RelationFollowerList(req.Token, user_id); err == nil {
		var userList []User
		for _, tmp := range respClient.UserList {
			userList = append(userList, *RPCUser2ControllerUser(tmp))
		}
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			UserList: userList,
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{})
	}
}

// FriendList 获取朋友列表
func FriendList(c *gin.Context) {
	var req UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("FriendListRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserListResponse{Response: Response{StatusCode: 404, StatusMsg: "参数错误"}})
		return
	}

	user_id, _ := strconv.ParseInt(req.UserID, 10, 64)

	if respClient, err := client.RelationFriendList(req.Token, user_id); err == nil {
		var userList []User
		for _, tmp := range respClient.UserList {
			userList = append(userList, *RPCUser2ControllerUser(tmp))
		}
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			UserList: userList,
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{})
	}
}
