include "user.thrift"

namespace go feed

typedef i64 int64
typedef i32 int32

struct douyin_feed_request {
  1: optional int64 latest_time; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
  2: optional string token; // 可选参数，登录用户设置
}

struct douyin_feed_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
  3: list<Video> video_list; // 视频列表
  4: optional int64 next_time; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct Video {
  1: required int64 id; // 视频唯一标识
  2: optional user.User author; // 视频作者信息
  3: required string play_url; // 视频播放地址
  4: optional string cover_url; // 视频封面地址
  5: optional int64 favorite_count; // 视频的点赞总数
  6: optional int64 comment_count; // 视频的评论总数
  7: optional bool is_favorite; // true-已点赞，false-未点赞
  8: optional string title; // 视频标题
}

service FeedService {
    douyin_feed_response GetVideo(1: douyin_feed_request req);
}




