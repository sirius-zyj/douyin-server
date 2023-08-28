package client

import (
	"context"
	"douyin-server/config"
	"douyin-server/rpc/kitex_gen/relation"
	"douyin-server/rpc/kitex_gen/relation/relationservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelationClient() {
	// 服务发现
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	c, err := relationservice.NewClient(config.RelationServiceName,
		// client.WithHostPorts(config.RelationAddr),
		client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	relationClient = c
}

func RelationAction(token string, follow_id int64, action_type string) (resp *relation.DouyinRelationActionResponse, err error) {
	resp = new(relation.DouyinRelationActionResponse)

	resp, err = relationClient.RelationAction(context.Background(), &relation.DouyinRelationActionRequest{
		Token:      token,
		FollowId:   follow_id,
		ActionType: action_type,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func RelationFollowList(token string, user_id int64) (resp *relation.DouyinRelationFollowListResponse, err error) {
	resp = new(relation.DouyinRelationFollowListResponse)

	resp, err = relationClient.RelationFollowList(context.Background(), &relation.DouyinRelationFollowListRequest{
		Token:  token,
		UserId: user_id,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func RelationFollowerList(token string, user_id int64) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	resp = new(relation.DouyinRelationFollowerListResponse)

	resp, err = relationClient.RelationFollowerList(context.Background(), &relation.DouyinRelationFollowerListRequest{
		Token:  token,
		UserId: user_id,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func RelationFriendList(token string, user_id int64) (resp *relation.DouyinRelationFriendListResponse, err error) {
	resp = new(relation.DouyinRelationFriendListResponse)

	resp, err = relationClient.RelationFriendList(context.Background(), &relation.DouyinRelationFriendListRequest{
		Token:  token,
		UserId: user_id,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
