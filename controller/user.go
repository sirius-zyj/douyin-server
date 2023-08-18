package controller

import (
	"net/http"
	"strconv"

	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		ID:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if respClient, err := client.Register(username, password); err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			UserId:   respClient.UserId,
			Token:    respClient.Token,
		})
	} else {
		c.JSON(http.StatusExpectationFailed, UserLoginResponse{})
	}

}

// 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if respClient, err := client.Login(username, password); err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			UserId:   respClient.UserId,
			Token:    respClient.Token,
		})
	} else {
		c.JSON(http.StatusExpectationFailed, UserLoginResponse{})
	}
}

// 用户信息
func UserInfo(c *gin.Context) {
	Id := c.Query("user_id")
	id, _ := strconv.ParseInt(Id, 10, 64)

	if respClient, err := client.UserInfo(id); err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *RPCUser2ControlUser(respClient),
		})
	} else {
		c.JSON(http.StatusExpectationFailed, UserLoginResponse{})
	}
}
