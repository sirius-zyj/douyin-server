package client

import (
	"context"
	"douyin-server/config"
	"douyin-server/rpc/kitex_gen/user"
	"douyin-server/rpc/kitex_gen/user/userservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserClient() {
	// 服务发现
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	c, err := userservice.NewClient(config.UserServiceName,
		// client.WithHostPorts(config.UserAddr),
		client.WithResolver(r))

	if err != nil {
		log.Fatal(err)
	}
	userClient = c
}

func Register(username string, password string) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)
	resp, err = userClient.Register(context.Background(), &user.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}

func Login(username string, password string) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)
	resp, err = userClient.Login(context.Background(), &user.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}

func UserInfo(user_id int64, token *string) (resp *user.User, err error) {
	resp = new(user.User)
	resp, err = userClient.UserInfo(context.Background(), &user.DouyinUserInfoRequest{
		UserId: user_id,
		Token:  token,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}
