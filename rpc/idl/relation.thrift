include "user.thrift"

namespace go relation

typedef i32 int32
typedef i64 int64

struct douyin_relation_action_request {
  1: required string token ; // 用户鉴权token
  2: required int64 follow_id; // 对方用户id
  3: required string action_type; // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
  1: required int32 status_code ; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
}

struct douyin_relation_follow_list_request {
  1: required int64 user_id ; // 用户id
  2: required string token; // 用户鉴权token
}

struct douyin_relation_follow_list_response {
  1: required int32 status_code ; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
  3: list<user.User> user_list; // 用户信息列表
}

struct douyin_relation_follower_list_request {
  1: required int64 user_id ; // 用户id
  2: required string token; // 用户鉴权token
}

struct douyin_relation_follower_list_response {
  1: required int32 status_code ; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
  3: list<user.User> user_list; // 用户列表
}

service RelationService{
    douyin_relation_action_response RelationAction(1: douyin_relation_action_request req)
    douyin_relation_follow_list_response RelationFollowList(1: douyin_relation_follow_list_request req)
    douyin_relation_follower_list_response RelationFollowerList(1: douyin_relation_follower_list_request req)
}