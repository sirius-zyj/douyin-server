package dao

import (
	// "gorm.io/gorm"
	"errors"
	"log"
)

func (Duser) TableName() string {
	return "users"
}

// 通过用户名查找用户
func GetUsersByUserName(userName string) (Duser, error) {
	var userlist = make([]Duser, 0)
	var user Duser
	if err := db.Where("name = ?", userName).Find(&userlist).Error; err != nil {
		log.Println("查询错误")
		return user, errors.New("查询出错")
	}

	if len(userlist) == 0 {
		log.Println("没有匹配的数据")
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
func CreateUser(user Duser) int64 {
	log.Println("开始注册2.5 ")
	if err := db.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return -1
	}
	return user.ID
}
