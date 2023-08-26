namespace go message

typedef i32 int32
typedef i64 int64

struct douyin_relation_action_request {
  1: required string token; // 用户鉴权token
  2: required int64 to_user_id; // 对方用户id
  3: required string action_type; // 1-发送消息
  4: required string content; // 消息内容
}

struct douyin_relation_action_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
}

struct douyin_message_chat_request {
  1: required string token; // 用户鉴权token
  2: required int64 to_user_id; // 对方用户id
  3: required int64 pre_msg_time;//上次最新消息的时间（新增字段-apk更新中）
}

struct douyin_message_chat_response {
  1: required int32 status_code; // 状态码，0-成功，其他值-失败
  2: optional string status_msg; // 返回状态描述
  3: list<Message> message_list; // 消息列表
}

struct Message {
  1: required int64 id; // 消息id
  2: required int64 to_user_id; // 该消息接收者的id
  3: required int64 from_user_id; // 该消息发送者的id
  4: required string content; // 消息内容
  5: optional string create_time; // 消息创建时间
}

service MessageService{
    douyin_message_chat_response MessageChat(1: douyin_message_chat_request req)
    douyin_relation_action_response MessageAction(1: douyin_relation_action_request req)
}