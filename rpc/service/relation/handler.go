package main

import (
	"context"
	"douyin-server/database"
	"douyin-server/database/dao"
	relation "douyin-server/rpc/kitex_gen/relation"
	"log"
	"strconv"
	"strings"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

func setFollowActionResponse(resp *relation.DouyinRelationActionResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	resp = new(relation.DouyinRelationActionResponse)

	token, follow_id, action_type := req.Token, req.FollowId, req.ActionType

	index := strings.Index(token, "*")
	user_id, _ := strconv.ParseInt(token[index+1:], 10, 64)

	if fo, err := database.GetFollowData(user_id, follow_id); err == nil {
		//获取到的表数据ID为0时代表没有该条关注数据
		if fo.Id == 0 {
			fo.User_id = user_id
			fo.Follow_id = follow_id
			fo.Action_type = action_type
			err = database.InsertFollow(&fo)
			if err != nil {
				setFollowActionResponse(resp, 404, "关注失败")
			} else {
				dao.UpdateUser("id", user_id, "follow_count", 1)     //follow_count+1
				dao.UpdateUser("id", follow_id, "follower_count", 1) //用户follower_count+1
				setFollowActionResponse(resp, 0, "关注成功")
			}
		} else {
			if action_type != fo.Action_type {
				if err := dao.EraseFollow(user_id, follow_id); err != nil {
					setFollowActionResponse(resp, 404, "关注数据erase失败")
				} else {
					dao.UpdateUser("id", user_id, "follow_count", -1)     //follow_count-1
					dao.UpdateUser("id", follow_id, "follower_count", -1) //用户follower_count-1
					setFollowActionResponse(resp, 0, "关注数据erase成功")
				}
			} else {
				if action_type != "1" {
					dao.EraseFollow(user_id, follow_id)
				}
				setFollowActionResponse(resp, 0, "Action_type 与数据库中的数据相同")
			}
		}
	} else {
		setFollowActionResponse(resp, 404, "关注发生错误")
	}
	return
}

func setFollowListResponse(resp *relation.DouyinRelationFollowListResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	resp = new(relation.DouyinRelationFollowListResponse)

	userID := req.UserId
	FollowIDList, err := dao.GetFollowList(userID)
	if err != nil {
		setFollowListResponse(resp, 404, "Follow列表查询失败")
	} else {
		var UserList []dao.Duser
		UserList, err = dao.GetUsersByIds(FollowIDList)
		if err != nil {
			log.Println("查询失败")
			setFollowListResponse(resp, 404, "Follow数据查询错误")
			return
		}
		setFollowListResponse(resp, 0, "Follow数据查询成功")
		for _, tmp := range UserList {
			resp.UserList = append(resp.UserList, database.DaoUser2RPCUser(&req.Token, &tmp))
		}
	}
	return
}

func setFollowerListResponse(resp *relation.DouyinRelationFollowerListResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	resp = new(relation.DouyinRelationFollowerListResponse)

	followID := req.UserId
	UserIDList, err := dao.GetFollowerList(followID)
	if err != nil {
		setFollowerListResponse(resp, 404, "Follower列表查询失败")
	} else {
		var UserList []dao.Duser
		UserList, err = dao.GetUsersByIds(UserIDList)
		if err != nil {
			log.Println("查询失败")
			setFollowerListResponse(resp, 404, "Follower数据查询错误")
			return
		}
		setFollowerListResponse(resp, 0, "Follower数据查询成功")
		for _, tmp := range UserList {
			resp.UserList = append(resp.UserList, database.DaoUser2RPCUser(&req.Token, &tmp))
		}
	}
	return
}

func setFriendListResponse(resp *relation.DouyinRelationFriendListResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	resp = new(relation.DouyinRelationFriendListResponse)

	FriendIDList, err := dao.GetFriendList(req.UserId)
	if err != nil {
		setFriendListResponse(resp, 404, "Friend列表查询失败")
	} else {
		var UserList []dao.Duser
		UserList, err = dao.GetUsersByIds(FriendIDList)
		if err != nil {
			log.Println("查询失败")
			setFriendListResponse(resp, 404, "Friend数据查询错误")
			return
		}
		setFriendListResponse(resp, 0, "Friend数据查询成功")
		for _, tmp := range UserList {
			resp.UserList = append(resp.UserList, database.DaoUser2RPCUser(&req.Token, &tmp))
		}
	}
	return
}
