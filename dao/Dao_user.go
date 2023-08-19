package dao

import (
	// "gorm.io/gorm"
	"errors"
	"log"

	"gorm.io/gorm"
)

func (Duser) TableName() string {
	return "users"
}

// 通过用户名查找用户
func GetUsersByUserName(userName string) (Duser, error) {
	var userlist = make([]Duser, 0)
	var user Duser
	if err := db.Where("name = ?", userName).Find(&userlist).Error; err != nil {
		return user, errors.New("查询出错")
	}

	if len(userlist) == 0 {
		return user, errors.New("没有匹配的用户")
	}
	user = userlist[0]
	return user, nil
}

// 通过用户id查找用户
func GetUserById(userid int64) (Duser, error) {
	var user Duser
	err := db.Where("id = ?", userid).Find(&user).Error
	if err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

// 通过用户id查找一系列用户
func GetUsersByIds(userid []int64) ([]Duser, error) {
	var resp []Duser
	err := db.Where("id = ?", userid).Find(&resp).Error
	if err != nil {
		log.Println(err.Error())
		return resp, err
	}
	return resp, nil
}

// 创建用户
func CreateUser(user *Duser) (err error) {
	if err := db.Create(user).Error; err != nil {
		log.Println("CreateUser Err : ", err.Error())
		return err
	}
	if err := db.Where("name = ?", user.Name).Find(user).Error; err != nil {
		log.Println("CreateUser Err : ", err.Error())
		return err
	}
	return
}

func UpdateUser(where string, where_count int64, what string, what_count int64) error {
	err := db.Model(&Duser{}).Where(where+" = ?", where_count).Update(what, gorm.Expr(what+" + ?", what_count)).Error
	if err != nil {
		log.Println("更新失败")
		return err
	}
	return nil
}

// func UpdateWorkCount(user_id, count int64) (err error) {
// 	if err := db.Model(&Duser{}).Where("id = ?", user_id).Update("work_count", gorm.Expr("work_count + ?", count)).Error; err != nil {
// 		log.Println("UpdateWorkCount Err : ", err.Error())
// 		return err
// 	}
// 	return
// }

// func UpdateUserFavoriteCount(user_id, count int64) (err error) {
// 	if err := db.Model(&Duser{}).Where("id = ?", user_id).Update("favorite_count", gorm.Expr("favorite_count + ?", count)).Error; err != nil {
// 		log.Println("UpdateFavoriteCount Err : ", err.Error())
// 		return err
// 	}
// 	return
// }
