package client

import (
	"context"
	"douyin-server/rpc/kitex_gen/user"
	"douyin-server/rpc/kitex_gen/user/userservice"
	"log"

	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func initUserClient() {
	c, err := userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8881"))
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

func UserInfo(user_id int64) (resp *user.User, err error) {
	resp = new(user.User)
	resp, err = userClient.UserInfo(context.Background(), user_id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}
