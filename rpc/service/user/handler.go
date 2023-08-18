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

func setRegisterResp(resp *user.DouyinUserRegisterResponse, statusCode int32, statusMsg string, userId int64, token string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
	resp.UserId = userId
	resp.Token = token
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)
	username, password := req.Username, req.Password
	//验证用户名是否已存在
	if _, err = dao.GetUsersByUserName(username); err == nil {
		//该用户名已存在
		setRegisterResp(resp, 404, "User already exist", -1, "")
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
			setRegisterResp(resp, 404, "User register failed", -1, "")
			return resp, nil
		}
		log.Println("注册成功")
		newUser.ID = ID
		token := username + password + "*" + strconv.FormatInt(ID, 10)
		setRegisterResp(resp, 0, "User register success", ID, token)
		return resp, nil
	}

}

func setLoginResp(resp *user.DouyinUserLoginResponse, statusCode int32, statusMsg string, userId int64, token string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
	resp.UserId = userId
	resp.Token = token
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)
	username, password := req.Username, req.Password
	if user, err := dao.GetUsersByUserName(username); err == nil {
		//找到了用户信息
		if password == user.Password {
			setLoginResp(resp, 0, "User login success", user.ID, username+password+"*"+strconv.FormatInt(user.ID, 10))
		} else {
			setLoginResp(resp, 404, "User password error", -1, "")
		}
	} else {
		setLoginResp(resp, 404, "User not exist", -1, "")
	}
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.DouyinUserInfoRequest) (resp *user.User, err error) {
	resp = new(user.User)
	if tmp, err := dao.GetUserById(req.UserId); err == nil {
		resp = &user.User{
			Id:              tmp.ID,
			Name:            tmp.Name,
			FollowCount:     &tmp.FollowCount,
			FollowerCount:   &tmp.FollowerCount,
			Avatar:          &tmp.Avatar,
			BackgroundImage: &tmp.BackgroundImage,
			Signature:       &tmp.Signature,
			TotalFavorited:  &tmp.TotalFavorited,
			WorkCount:       &tmp.WorkCount,
			FavoriteCount:   &tmp.FavoriteCount,
			// IsFollow: true, // TODO
		}
		return resp, nil
	}
	return resp, err
}
