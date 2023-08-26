package dao

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func (Dmessage) TableName() string {
	return "messages"
}

func InsertMessage(message *Dmessage) (err error) {
	if err = db.Model(&Dmessage{}).Create(message).Error; err != nil {
		return err
	}
	return nil
}

func GetMessageList(fromUserId int64, toUserId int64, preMsgTime time.Time) (messageList []Dmessage, err error) {
	if err = db.Model(&Dmessage{}).
		Where("((from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)) AND created_at > ?", fromUserId, toUserId, toUserId, fromUserId, preMsgTime).
		Order("created_at ASC").
		Find(&messageList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return messageList, nil
		}
		return nil, err
	}
	return
}
