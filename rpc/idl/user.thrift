namespace go user

typedef i64 int64
typedef i32 int32

struct douyin_user_register_request {
  1: required string username; // 注册用户名，最长32个字符
  2: required string password; // 密码，最长32个字符
}

struct douyin_user_register_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败 404-用户名已存在
  2: optional string status_msg; // 返回状态描述
  3: required int64 user_id; // 用户id
  4: required string token; // 用户鉴权token
}

struct douyin_user_login_request {
  1: required string username; // 登录用户名
  2: required string password; // 登录密码
}

struct douyin_user_login_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
  3: required int64 user_id; // 用户id
  4: required string token; // 用户鉴权token
}

struct douyin_user_info_request {
  1: required int64 user_id; // 用户id
  2: optional string token; // 用户鉴权token
}

struct User {
  1: required int64 id; // 用户id
  2: required string name; // 用户名称
  3: optional int64 follow_count; // 关注总数
  4: optional int64 follower_count; // 粉丝总数
  5: required bool is_follow; // true-已关注，false-未关注
  6: optional string avatar; //用户头像
  7: optional string background_image; //用户个人页顶部大图
  8: optional string signature; //个人简介
  9: optional int64 total_favorited; //获赞数量
  10: optional int64 work_count; //作品数量
  11: optional int64 favorite_count; //点赞数量
}

service UserService {
    douyin_user_register_response Register(1: douyin_user_register_request req);
    douyin_user_login_response Login(1: douyin_user_login_request req);
    User UserInfo(1: douyin_user_info_request req);
}