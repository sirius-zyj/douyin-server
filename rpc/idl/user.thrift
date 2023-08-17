namespace go user

typedef i64 int64
typedef i32 int32

message douyin_user_register_request {
  required string username = 1; // 注册用户名，最长32个字符
  required string password = 2; // 密码，最长32个字符
}

message douyin_user_register_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  required int64 user_id = 3; // 用户id
  required string token = 4; // 用户鉴权token
}

message douyin_user_login_request {
  required string username = 1; // 登录用户名
  required string password = 2; // 登录密码
}

message douyin_user_login_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  required int64 user_id = 3; // 用户id
  required string token = 4; // 用户鉴权token
}