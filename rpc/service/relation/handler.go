package main

import (
	"context"
	"douyin-server/dao"
	relation "douyin-server/rpc/kitex_gen/relation"
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

	fo, err := dao.GetFollowData(user_id, follow_id)
	if err == nil {
		//获取到的表数据ID为0时代表没有该条关注数据
		if fo.Id == 0 {
			fo.User_id = user_id
			fo.Follow_id = follow_id
			fo.Action_type = action_type
			err = dao.InsertFollow(fo)
			if err != nil {
				setFollowActionResponse(resp, 404, "关注失败")
			} else {
				dao.UpdateUser("id", user_id, "follow_count", 1)     //follow_count+1
				dao.UpdateUser("id", follow_id, "follower_count", 1) //用户follower_count+1
				setFollowActionResponse(resp, 0, "关注成功")
			}
		} else {
			if fo.Action_type != action_type {
				err := dao.ActionFollow(user_id, follow_id, action_type)
				if err != nil {
					setFollowActionResponse(resp, 404, "关注数据更新失败")
				} else {
					if action_type == "1" {
						dao.UpdateUser("id", user_id, "follow_count", 1)     //follow_count+1
						dao.UpdateUser("id", follow_id, "follower_count", 1) //用户follower_count+1
					} else {
						dao.UpdateUser("id", user_id, "follow_count", -1)     //follow_count-1
						dao.UpdateUser("id", follow_id, "follower_count", -1) //用户follower_count-1
					}
					setFollowActionResponse(resp, 0, "关注数据更新成功")
				}
			} else {
				setFollowActionResponse(resp, 0, "Action_type 与数据库中的数据相同")
			}
		}
	} else {
		setFollowActionResponse(resp, 404, "关注发生错误")
	}
	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
