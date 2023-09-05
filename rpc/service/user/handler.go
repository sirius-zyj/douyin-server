package main

import (
	"context"
	"douyin-server/database"
	"douyin-server/database/dao"
	"douyin-server/middleware/jwt"
	user "douyin-server/rpc/kitex_gen/user"
	"log"
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
	username, password := req.Username, jwt.EnCoder(req.Password)
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
		if err := dao.CreateUser(&newUser); err != nil {
			log.Println("注册失败")
			setRegisterResp(resp, 404, "User register failed", -1, "")
			return resp, nil
		}

		// token := username + password + "*" + strconv.FormatInt(newUser.ID, 10)
		token := jwt.GenerateToken(username)

		setRegisterResp(resp, 0, "User register success", newUser.ID, token)
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
		if jwt.EnCoder(password) == user.Password {
			token := jwt.GenerateToken(username)
			setLoginResp(resp, 0, "User login success", user.ID, token)
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

	// if tmp, err := dao.GetUserById(req.UserId); err == nil {
	if tmp, err := database.GetUserById(req.UserId); err == nil {
		resp = database.DaoUser2RPCUser(req.Token, &tmp)
		return resp, nil
	}
	return resp, err
}
