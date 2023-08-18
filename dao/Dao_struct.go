package dao

import (
	"time"
)

type Dvideo struct {
	Id             int64 `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Author_id      int64
	Play_url       string `json:"play_url"`
	Cover_url      string `json:"cover_url"`
	Upload_time    time.Time
	Title          string `json:"title"`
	Favorite_count int64  `json:"favorite_count"`
	Comment_count  int64  `json:"comment_count"`
}

type Duser struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	Password        string `json:"password"`         // 用户密码
	// IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name           string `json:"name"`            // 用户名称
	Signature      string `json:"signature"`       // 个人简介
	TotalFavorited int64  `json:"total_favorited"` // 获赞数量
	WorkCount      int64  `json:"work_count"`      // 作品数
}

type Dfavorite struct {
	Id       int64 `gorm:"primaryKey;autoIncrement"`
	User_id  int64
	Video_id int64
	Cancel   bool
}

type Dcomments struct {
	Id           int64 `gorm:"primaryKey;autoIncrement"`
	User_id      int64
	Video_id     int64
	Comment_text string
	Created_at   time.Time
}
