package dao

import (
	"time"
)

type Dvideo struct {
	Id             int64     `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Author_id      int64     `json:"author_id"`
	Play_url       string    `json:"play_url"`
	Cover_url      string    `json:"cover_url"`
	Upload_time    time.Time `json:"upload_time"`
	Title          string    `json:"title"`
	Favorite_count int64     `json:"favorite_count"`
	Comment_count  int64     `json:"comment_count"`
}

type Duser struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	Password        string `json:"password"`         // 用户密码
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数量
}

type Dfavorite struct {
	Id          int64  `gorm:"primaryKey;autoIncrement"`
	User_id     int64  `json:"user_id"`
	Video_id    int64  `json:"video_id"`
	Action_type string `json:"action_type"`
}

type Dcomments struct {
	Id           int64     `gorm:"primaryKey;autoIncrement"`
	User_id      int64     `json:"user_id"`
	Video_id     int64     `json:"video_id"`
	Comment_text string    `json:"comment_text"`
	Created_at   time.Time `json:"created_at"`
}

type Dfollow struct {
	Id          int64  `gorm:"primaryKey;autoIncrement"`
	User_id     int64  `json:"user_id"`
	Follow_id   int64  `json:"follow_id"`
	Action_type string `json:"action_type"`
}

type Dmessage struct {
	Id           int64     `gorm:"primaryKey;autoIncrement"`
	To_user_id   int64     `json:"to_user_id"`
	From_user_id int64     `json:"from_user_id"`
	Content      string    `json:"content"`
	Created_at   time.Time `json:"created_at"`
}
