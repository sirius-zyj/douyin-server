package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
  "github.com/RaymondCode/simple-demo/dao"
  "strconv"
  "strings"
  "time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction 发表或者删除评论
func CommentAction(c *gin.Context) {
	token := c.Query("token")
  index := strings.Index(token , "*")
  useid := token[index + 1:]
  user_id , _ := strconv.ParseInt(useid , 10 , 64)
	actionType := c.Query("action_type")
  if actionType == "1" {
    var comment dao.Dcomments
    comment.Comment_text = c.Query("comment_text")
    comment.User_id = user_id
    comment.Video_id , _ = strconv.ParseInt(c.Query("video_id") , 10 , 64)
    comment.Create_time = time.Now()
  }

  
	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList 查看评论列表
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
