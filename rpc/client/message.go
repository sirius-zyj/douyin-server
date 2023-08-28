package client

import (
	"context"
	"douyin-server/config"
	"douyin-server/rpc/kitex_gen/message"
	"douyin-server/rpc/kitex_gen/message/messageservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func initMessageClient() {
	// 服务发现
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	c, err := messageservice.NewClient(config.MessageServiceName,
		// client.WithHostPorts(config.MessageAddr),
		client.WithResolver(r))
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
