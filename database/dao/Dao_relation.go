package dao

import (
	"database/sql"
	"errors"
	"log"

	"gorm.io/gorm"
)

func (Dfollow) TableName() string {
	return "follows"
}

// 关注erase操作  根据关注用户和被关注用户id进行访问
func EraseFollow(user_id int64, follower_id int64) error {
	err := db.Where("user_id = ? AND follow_id = ?", user_id, follower_id).Delete(&Dfollow{}).Error

	if err != nil {
		log.Println("关注更新错误")
		return errors.New("关注更新错误")
	}
	return nil
}

// 插入关注数据
func InsertFollow(foDate *Dfollow) error {
	err := db.Model(&Dfollow{}).Create(foDate).Error
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

// 查询某用户的关注ID列表
func GetFollowList(user_id int64) (resp []int64, err error) {
	var res []sql.NullInt64
	err = db.Model(&Dfollow{}).Where("user_id = ?", user_id).Pluck("follow_id", &res).Error
	for _, tmp := range res {
		if tmp.Valid {
			resp = append(resp, tmp.Int64)
		}
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 查询为空
			log.Println("查询为空")
			return resp, nil
		} else {
			// 其他错误
			log.Println("关注数据查询发生错误")
			return resp, errors.New("关注数据查询发生错误")
		}
	}
	return resp, nil
}

// 查询某用户的follower ID列表
func GetFollowerList(follow_id int64) (resp []int64, err error) {
	var res []sql.NullInt64
	err = db.Model(&Dfollow{}).Where("follow_id = ?", follow_id).Pluck("user_id", &res).Error
	for _, tmp := range res {
		if tmp.Valid {
			resp = append(resp, tmp.Int64)
		}
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 查询为空
			log.Println("查询为空")
			return resp, nil
		} else {
			// 其他错误
			log.Println("关注数据查询发生错误")
			return resp, errors.New("关注数据查询发生错误")
		}
	}
	return resp, nil
}

func GetFriendList(user_id int64) (resp []int64, err error) {
	var res []sql.NullInt64
	if err = db.Model(&Dfollow{}).Where("user_id = ?", user_id).Pluck("follow_id", &res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 查询为空
			log.Println("查询为空")
			return resp, nil
		} else {
			// 其他错误
			log.Println("关注数据查询发生错误")
			return resp, errors.New("关注数据查询发生错误")
		}
	} else {
		for _, tmp := range res {
			if tmp.Valid {
				if err = db.Model(&Dfollow{}).Where("user_id = ? And follow_id = ?", tmp.Int64, user_id).Error; err == nil {
					resp = append(resp, tmp.Int64)
				}
			}
		}
	}
	return resp, nil
}
