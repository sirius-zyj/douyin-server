package controller

import (
	"douyin-server/rpc/client"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var tempChat = map[string][]Message{}

var messageIdSequence = int64(1)

type MessageActionRequest struct {
	ActionType string `form:"action_type"` // 1-发送消息
	Content    string `form:"content"`     // 消息内容
	ToUserID   string `form:"to_user_id"`  // 对方用户id
	Token      string `form:"token"`       // 用户鉴权token
}

type MessageChatRequest struct {
	ToUserID   string `form:"to_user_id"`   // 对方用户id
	Token      string `form:"token"`        // 用户鉴权token
	PreMsgTime int64  `form:"pre_msg_time"` // 上次获取消息的时间
}

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	var req MessageActionRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("MessageActionRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserResponse{Response: Response{StatusCode: 404}})
		return
	}

	toUserId, _ := strconv.ParseInt(req.ToUserID, 10, 64)
	if respClient, err := client.MessageAction(req.Token, toUserId, req.ActionType, req.Content); err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)})
	} else {
		c.JSON(http.StatusInternalServerError, Response{})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	var req MessageChatRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("MessageChatRequest Err : ", err)
		c.JSON(http.StatusBadRequest, UserResponse{Response: Response{StatusCode: 404}})
		return
	}

	if req.PreMsgTime == 0 {
		c.JSON(http.StatusOK, ChatResponse{Response{StatusCode: 0, StatusMsg: "no new messages"}, []Message{}})
		return
	}

	toUserId, _ := strconv.ParseInt(req.ToUserID, 10, 64)
	if respClient, err := client.MessageChat(req.Token, toUserId, req.PreMsgTime); err == nil {
		var messageList []Message
		for tmpId, tmp := range respClient.MessageList {
			// createTime, _ := strconv.ParseInt(*tmp.CreateTime, 10, 64)
			messageList = append(messageList, Message{
				Id:         int64(tmpId),
				FromUserId: tmp.FromUserId,
				ToUserId:   tmp.ToUserId,
				Content:    tmp.Content,
				CreateTime: 0,
			})
		}
		c.JSON(http.StatusOK, ChatResponse{
			Response:    Response{StatusCode: respClient.StatusCode, StatusMsg: StatusMsg(respClient.StatusMsg)},
			MessageList: messageList,
		})
	} else {
		c.JSON(http.StatusInternalServerError, ChatResponse{})
	}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
