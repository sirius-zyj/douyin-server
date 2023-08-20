package dao

import (
	"errors"
	"log"
)

func (Dfollow) TableName() string {
	return "follows"
}

// 关注更新操作  根据关注用户和被关注用户id进行访问
func ActionFollow(user_id int64, follower_id int64, action_type string) error {
	err := db.Model(&Dfollow{}).Where("user_id = ? AND follow_id = ?", user_id, follower_id).Update("action_type", action_type).Error

	if err != nil {
		log.Println("关注更新错误")
		return errors.New("关注更新错误")
	}
	return nil
}

// 插入关注数据
func InsertFollow(foDate Dfollow) error {
	err := db.Model(&Dfollow{}).Create(&foDate).Error
	if err != nil {
		log.Println("关注数据插入失败")
		return errors.New("关注数据插入失败")
	}
	return nil
}

// 查询关注信息
func GetFollowData(user_id int64, follow_id int64) (Dfollow, error) {
	var followData Dfollow
	err := db.Model(&Dfollow{}).Where("user_id = ? AND follow_id = ?", user_id, follow_id).Find(&followData).Error
	if err != nil {
		log.Println("查询错误")
		return followData, errors.New("查询出错")
	}
	return followData, nil
}
