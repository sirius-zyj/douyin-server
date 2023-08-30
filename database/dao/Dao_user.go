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
	if err := db.Model(&Duser{}).Where("name = ?", userName).Find(&userlist).Error; err != nil {
		return user, errors.New("查询出错")
	}
	log.Println(userlist)
	if len(userlist) == 0 {
		return user, errors.New("没有匹配的用户")
	}
	user = userlist[0]
	return user, nil
}

// 通过用户id查找用户
func GetUserById(userid int64) (Duser, error) {
	var user Duser
	err := db.Model(&Duser{}).Where("id = ?", userid).Find(&user).Error
	if err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

// 通过用户id查找一系列用户
func GetUsersByIds(userids []int64) ([]Duser, error) {
	var resp []Duser
	err := db.Model(&Duser{}).Where("id in ?", userids).Find(&resp).Error
	if err != nil {
		log.Println(err.Error())
		return resp, err
	}
	return resp, nil
}

// 创建用户
func CreateUser(user *Duser) (err error) {
	if err := db.Model(&Duser{}).Create(user).Error; err != nil {
		log.Println("CreateUser Err : ", err.Error())
		return err
	}
	if err := db.Where("name = ?", user.Name).Find(user).Error; err != nil {
		log.Println("CreateUser Err : ", err.Error())
		return err
	}
	return
}
