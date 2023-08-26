package dao

import (
	"errors"
	"log"
)

func (Dcomments) TableName() string {
	return "comments"
}

func InsertComment(comment Dcomments) error {
	err := db.Model(&Dcomments{}).Create(&comment).Error
	if err != nil {
		log.Println("创建评论失败")
		return err
	}
	return nil
}

func EraseComment(comment_id, video_id int64) error {
	err := db.Model(&Dcomments{}).Where("id = ? AND video_id = ?", comment_id, video_id).Delete(&Dcomments{}).Error
	if err != nil {
		log.Println("评论删除失败")
		return err
	}
	return nil
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
