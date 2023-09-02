package controller

import (
	"log"
	"net/http"
	"strconv"

	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	Password string `form:"password"` // 密码，最长32个字符
	Username string `form:"username"` // 注册用户名，最长32个字符
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserRequest struct {
	Token  *string `form:"token"`   // 用户鉴权token
	UserID string  `form:"user_id"` // 用户id
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// 用户注册
func Register(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("UserRegisterRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserLoginResponse{Response: Response{StatusCode: 404, StatusMsg: "参数错误"}})
		return
	}

	if respClient, err := client.Register(req.Username, req.Password); err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			UserId:   respClient.UserId,
			Token:    respClient.Token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{})
	}

}

// 用户登录
func Login(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("UserLoginRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserLoginResponse{Response: Response{StatusCode: 404}})
		return
	}

	if respClient, err := client.Login(req.Username, req.Password); err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			UserId:   respClient.UserId,
			Token:    respClient.Token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{Response: Response{StatusCode: 404}})
	}
}

// 用户信息
func UserInfo(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("UserInfoRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserResponse{Response: Response{StatusCode: 404}})
		return
	}
	id, _ := strconv.ParseInt(req.UserID, 10, 64)
	//TODO token is unknown used
	if respClient, err := client.UserInfo(id, req.Token); err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "success"},
			User:     *RPCUser2ControllerUser(respClient),
		})
	} else {
		c.JSON(http.StatusInternalServerError, UserResponse{Response: Response{StatusCode: 404}})
	}
}
