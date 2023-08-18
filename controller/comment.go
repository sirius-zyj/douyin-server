package controller

import (
	"log"
	"net/http"
	"strconv"

	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
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
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	commentText := c.Query("comment_text")
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	if respClient, err := client.CommentAction(token, videoId, int32(actionType), &commentText, commentId); err == nil {
		log.Println("CommentAction:  ", *respClient)
		if respClient.Comment != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
				Comment: Comment{
					Id:         respClient.Comment.Id,
					User:       User{Id: respClient.Comment.User.Id, Name: respClient.Comment.User.Name},
					Content:    respClient.Comment.Content,
					CreateDate: respClient.Comment.CreateDate,
				},
			})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			})
		}
	} else {
		c.JSON(http.StatusExpectationFailed, CommentActionResponse{})
	}
}

// CommentList 查看评论列表
func CommentList(c *gin.Context) {
	videoid := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoid, 10, 64)

	if respClient, err := client.CommentList("", videoId); err == nil {
		log.Println("CommentList:  ", *respClient)
		var CommentResq []Comment
		for _, tmp := range respClient.CommentList {
			CommentResq = append(CommentResq, Comment{
				Id:         tmp.Id,
				User:       User{Id: tmp.User.Id, Name: tmp.User.Name},
				Content:    tmp.Content,
				CreateDate: tmp.CreateDate,
			})
		}
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			CommentList: CommentResq,
		})
	} else {
		c.JSON(http.StatusExpectationFailed, CommentListResponse{})
	}
}
