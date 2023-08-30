package dao

import (
	"errors"
	"log"
)

func (Dcomments) TableName() string {
	return "comments"
}

func GetAllComments(videoId int64) ([]Dcomments, error) {
	var CommentsList []Dcomments
	err := db.Model(&Dcomments{}).Where("video_id = ?", videoId).Find(&CommentsList).Error
	if err != nil {
		log.Println("评论拉取失败")
		return CommentsList, errors.New("评论拉取失败")
	}
	return CommentsList, nil
}
