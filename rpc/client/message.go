package client

import (
	"context"
	"douyin-server/rpc/kitex_gen/message"
	"douyin-server/rpc/kitex_gen/message/messageservice"
	"log"

	"github.com/cloudwego/kitex/client"
)

var messageClient messageservice.Client

func initMessageClient() {
	c, err := messageservice.NewClient("MessageService", client.WithHostPorts("0.0.0.0:8886"))
	if err != nil {
		log.Fatal(err)
	}
	messageClient = c
}

func MessageAction(token string, toUserId int64, action_type string, content string) (resp *message.DouyinRelationActionResponse, err error) {
	resp = new(message.DouyinRelationActionResponse)
	if resp, err = messageClient.MessageAction(context.Background(), &message.DouyinRelationActionRequest{
		Token:      token,
		ToUserId:   toUserId,
		ActionType: action_type,
		Content:    content,
	}); err != nil {
		log.Println("MessageAction Err : ", err)
		return nil, err
	}
	return
}

func MessageChat(token string, toUserId int64, preMsgTime int64) (resp *message.DouyinMessageChatResponse, err error) {
	resp = new(message.DouyinMessageChatResponse)
	if resp, err = messageClient.MessageChat(context.Background(), &message.DouyinMessageChatRequest{
		Token:      token,
		ToUserId:   toUserId,
		PreMsgTime: preMsgTime,
	}); err != nil {
		log.Println("MessageChat Err : ", err)
		return nil, err
	}
	return
}
