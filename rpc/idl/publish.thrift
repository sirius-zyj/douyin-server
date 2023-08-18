include "feed.thrift"

namespace go publish

typedef i32 int32
typedef i64 int64

struct douyin_publish_action_request {
  1: required string token; // 用户鉴权token
  2: required binary data; // 视频数据
  3: required string title; // 视频标题
}

struct douyin_publish_action_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
}

struct douyin_publish_list_request {
  1: required int64 user_id; // 用户id
  2: required string token; // 用户鉴权token
}

struct douyin_publish_list_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
  3: list<feed.Video> video_list; // 用户发布的视频列表
}

service PublishService {
    douyin_publish_action_response Publish(1: douyin_publish_action_request req)
    douyin_publish_list_response PublishList(1: douyin_publish_list_request req)
}