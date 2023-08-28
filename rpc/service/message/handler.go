package main

import (
	"context"
	"douyin-server/database/dao"
	"douyin-server/middleware/jwt"
	message "douyin-server/rpc/kitex_gen/message"
	"time"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

func setRelationActionResponse(resp *message.DouyinRelationActionResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// Action implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinRelationActionRequest) (resp *message.DouyinRelationActionResponse, err error) {
	resp = new(message.DouyinRelationActionResponse)

	if req.Content == "" {
		setRelationActionResponse(resp, 404, "发送内容不能为空")
		return
	}

	fromUserId := jwt.GetUserIdByToken(req.Token)

	if err = dao.InsertMessage(&dao.Dmessage{
		From_user_id: fromUserId,
		To_user_id:   req.ToUserId,
		Content:      req.Content,
		Created_at:   time.Now(),
	}); err != nil {
		setRelationActionResponse(resp, 404, "发送失败")
	} else {
		setRelationActionResponse(resp, 0, "发送成功")
	}
	return
}

func setMessageChatResponse(resp *message.DouyinMessageChatResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	resp = new(message.DouyinMessageChatResponse)

	fromUserId := jwt.GetUserIdByToken(req.Token)

	seconds := req.PreMsgTime / 1000
	nanoseconds := (req.PreMsgTime % 1000) * 1000000
	preMsgTime := time.Unix(seconds, nanoseconds)

	if messageList, err := dao.GetMessageList(fromUserId, req.ToUserId, preMsgTime); err != nil {
		setMessageChatResponse(resp, 404, "获取失败")
	} else {
		setMessageChatResponse(resp, 0, "获取成功")
		for _, tmp := range messageList {
			messageTime := tmp.Created_at.Format("2006-01-02 15:04:05")
			resp.MessageList = append(resp.MessageList, &message.Message{
				Id:         tmp.Id,
				FromUserId: tmp.From_user_id,
				ToUserId:   tmp.To_user_id,
				Content:    tmp.Content,
				CreateTime: &messageTime,
			})
		}
	}
	return
}
