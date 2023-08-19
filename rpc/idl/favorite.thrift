include "feed.thrift"

namespace go favorite

typedef i32 int32
typedef i64 int64

struct douyin_favorite_action_request {
  1: required string token; // 用户鉴权token
  2: required int64 video_id; // 视频id
  3: required string action_type; // 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
}

struct douyin_favorite_list_request {
  1: required int64 user_id; // 用户id
  2: required string token; // 用户鉴权token
}

struct douyin_favorite_list_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
  3: list<feed.Video> video_list; // 用户点赞视频列表
}

service FavoriteService {
    douyin_favorite_action_response FavoriteAction(1: douyin_favorite_action_request req);
    douyin_favorite_list_response FavoriteList(1: douyin_favorite_list_request req);
}

