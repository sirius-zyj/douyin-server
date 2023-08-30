package controller

import (
	"log"
	"net/http"
	"strconv"

	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

type CommentActionRequest struct {
	ActionType  string  `form:"action_type"`            // 1-发布评论，2-删除评论
	CommentID   *string `form:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
	CommentText *string `form:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	Token       string  `form:"token"`                  // 用户鉴权token
	VideoID     string  `form:"video_id"`               // 视频id
}

type CommentActionResponse struct {
	Response
	Comment *Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentAction 发表或者删除评论
func CommentAction(c *gin.Context) {
	var req CommentActionRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("CommentActionRequest Err : ", err)
		c.JSON(http.StatusBadRequest, CommentActionResponse{Response: Response{StatusCode: 404}})
		return
	}

	videoId, _ := strconv.ParseInt(req.VideoID, 10, 64)
	actionType, _ := strconv.ParseInt(req.ActionType, 10, 64)

	if respClient, err := client.CommentAction(req.Token, videoId, int32(actionType), req.CommentText, req.CommentID); err == nil {
		if respClient.Comment != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
				Comment: &Comment{
					Id:         respClient.Comment.Id,
					User:       *RPCUser2ControllerUser(respClient.Comment.User),
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
	c.JSON(http.StatusOK, CommentListResponse{})
	// videoid := c.Query("video_id")
	// videoId, _ := strconv.ParseInt(videoid, 10, 64)
	// token := c.Query("token")

	// if respClient, err := client.CommentList(token, videoId); err == nil {
	// 	var CommentResq []Comment
	// 	for _, tmp := range respClient.CommentList {
	// 		CommentResq = append(CommentResq, Comment{
	// 			Id:         tmp.Id,
	// 			User:       *RPCUser2ControllerUser(tmp.User),
	// 			Content:    tmp.Content,
	// 			CreateDate: tmp.CreateDate,
	// 		})
	// 	}
	// 	c.JSON(http.StatusOK, CommentListResponse{
	// 		Response:    Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
	// 		CommentList: CommentResq,
	// 	})
	// } else {
	// 	c.JSON(http.StatusExpectationFailed, CommentListResponse{})
	// }
}
