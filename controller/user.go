package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
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

//用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

  //验证用户名是否已存在
  if _, err := dao.GetUsersByUserName(username); err == nil {
    //该用户名已存在
    c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist1111"},
		})
  }else {
    //成功注册
    newUser := dao.Duser{
			Name: username,
      Password: password,
		}
    ID := dao.CreateUser(newUser)
    log.Println("开始注册3")
    if ID == -1 {
      log.Println("注册失败")
      c.JSON(http.StatusOK, UserLoginResponse{
		  	Response: Response{StatusCode: 1, StatusMsg: "User register failed1"},
		  })
      return 
    }
    newUser.Id = ID
    log.Println(err)
    token = token + "*" + strconv.FormatInt(ID , 10)
    c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   newUser.Id,
			Token:    token,
		})
  }
}

//用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password
  log.Println(token)
  if user , err := dao.GetUsersByUserName(username); err == nil {
    //找到了用户信息
    if(token == (user.Name + user.Password)) {
      c.JSON(http.StatusOK, UserLoginResponse{
  			Response: Response{StatusCode: 0},
  			UserId:   user.Id,
  			Token:    token,
		  })
    }else{
      c.JSON(http.StatusOK, UserLoginResponse{
  			Response: Response{StatusCode: 1 , StatusMsg: "Password error"},
		  })
    }
  }else{
    c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
  }
}

//用户信息
func UserInfo(c *gin.Context) {
	// token := c.Query("token")
  Id := c.Query("user_id")
  id , _ := strconv.ParseInt(Id , 10 , 64)
	if user , err := dao.GetUserById(id); err == nil {
    //找到了用户信息
    var resq User
    resq.Id = user.Id
    resq.Name = user.Name
    resq.FollowCount = user.FollowCount
    resq.FollowerCount = user.FollowerCount
    c.JSON(http.StatusOK, UserResponse{
      Response: Response{StatusCode: 0},
      User: resq,
    })
  }else{
    c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
  }
}

// func UserInfo(c *gin.Context) {
// 	token := c.Query("token")

// 	if user, exist := usersLoginInfo[token]; exist {
// 		c.JSON(http.StatusOK, UserResponse{
// 			Response: Response{StatusCode: 0},
// 			User:     user,
// 		})
// 	} else {
// 		c.JSON(http.StatusOK, UserResponse{
// 			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
// 		})
// 	}
// }
