package dao
import(
  "time"
)
type Dvideo struct{
  Id                 int64     `json:"id,omitempty"`
  Author_id          int64 
  Play_url           string    `json:"play_url"`
  Cover_url          string    `json:"cover_url"`
  Upload_time        time.Time 
  Title              string    `json:"title"`
  Favorite_count     int64     `json:"favorite_count"`
  Cosmment_count      int64    `json:"comment_count"`
}

type Duser struct{
  Id               int64
  Name             string
  Password         string
  FollowCount      int64 `gorm:"default:0" json:"following_count"`
  FollowerCount    int64 `gorm:"default:0" json:"follower_count"`
}