package main

import (
	"context"
	"douyin-server/dao"
	user "douyin-server/rpc/kitex_gen/user"
	"log"
	"strconv"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)
	username, password := req.Username, req.Password
	//验证用户名是否已存在
	if _, err = dao.GetUsersByUserName(username); err == nil {
		//该用户名已存在
		log.Println("该用户名已存在")
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "User already exist"
		resp.UserId = -1
		resp.Token = ""
		return resp, nil
	} else {
		//成功注册
		newUser := dao.Duser{
			Name:     username,
			Password: password,
		}
		ID := dao.CreateUser(newUser)
		if ID == -1 {
			log.Println("注册失败")
			resp.StatusCode = 404
			resp.StatusMsg = new(string)
			*resp.StatusMsg = "User register failed"
			resp.UserId = -1
			resp.Token = ""
			return resp, nil
		}
		log.Println("注册成功")
		newUser.Id = ID
		token := username + password + "*" + strconv.FormatInt(ID, 10)
		resp.StatusCode = 0
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "User register success"
		resp.UserId = ID
		resp.Token = token
		return resp, nil
	}

	return nil, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)
	username, password := req.Username, req.Password
	token := username + password
	if user, err := dao.GetUsersByUserName(username); err == nil {
		//找到了用户信息
		if token == (user.Name + user.Password) {
			resp.StatusCode = 0
			resp.StatusMsg = new(string)
			*resp.StatusMsg = "User login success"
			resp.UserId = user.Id
			resp.Token = token + "*" + strconv.FormatInt(user.Id, 10)
		} else {
			resp.StatusCode = 404
			resp.StatusMsg = new(string)
			*resp.StatusMsg = "User password error"
			resp.UserId = -1
			resp.Token = ""
		}
	} else {
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "User not exist"
		resp.UserId = -1
		resp.Token = ""
	}
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, userId user.Int64) (resp *user.User, err error) {
	resp = new(user.User)
	if user, err := dao.GetUserById(userId); err == nil {
		log.Println(user)
		//找到了用户信息
		resp.Id = user.Id
		resp.Name = user.Name
		resp.FollowCount = &user.FollowCount
		resp.FollowerCount = &user.FollowerCount
		return resp, nil
	}
	return resp, err
}
