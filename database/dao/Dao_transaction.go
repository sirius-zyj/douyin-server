package dao

import (
	"log"

	"gorm.io/gorm"
)

func Tran_InsertComment(comment Dcomments) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Dcomments{}).Create(&comment).Error; err != nil {
			log.Println("创建评论失败")
			return err
		}

		if err := tx.Model(&Dvideo{}).Where("id = ?", comment.Video_id).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			log.Println("创建评论, 更新失败")
			return err
		}

		return nil
	})
}

func Tran_EraseComment(comment_id, video_id int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Dcomments{}).Where("id = ? AND video_id = ?", comment_id, video_id).Delete(&Dcomments{}).Error; err != nil {
			log.Println("评论删除失败")
			return err
		}

		if err := tx.Model(&Dvideo{}).Where("id = ?", video_id).Update("comment_count", gorm.Expr("comment_count + ?", -1)).Error; err != nil {
			log.Println("评论删除, 更新失败")
			return err
		}

		return nil
	})
}

// 插入点赞数据
func Tran_InsertFavorite(faDate *Dfavorite, authorId int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Dfavorite{}).Create(faDate).Error; err != nil {
			log.Println("点赞数据插入失败")
			return err
		}

		if err := tx.Model(&Dvideo{}).Where("id = ?", faDate.Video_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			log.Println("点赞数据, 更新失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", authorId).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
			log.Println("点赞数据, 更新失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", faDate.User_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			log.Println("点赞数据, 更新失败")
			return err
		}

		return nil
	})
}

// Erase更新操作  根据点赞用户和视频id进行访问
func Tran_EraseFavorite(user_id, video_id, author_id int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Dfavorite{}).Where("user_id = ? AND video_id = ?", user_id, video_id).Delete(&Dfavorite{}).Error; err != nil {
			log.Println("点赞数据delete失败")
			return err
		}

		if err := tx.Model(&Dvideo{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count + ?", -1)).Error; err != nil {
			log.Println("点赞数据, 更新失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", author_id).Update("total_favorited", gorm.Expr("total_favorited + ?", -1)).Error; err != nil {
			log.Println("点赞数据, 更新失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", user_id).Update("favorite_count", gorm.Expr("favorite_count + ?", -1)).Error; err != nil {
			log.Println("点赞数据, 更新失败")
			return err
		}

		return nil
	})
}

func Tran_InsertVideo(video Dvideo) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Dvideo{}).Create(&video).Error; err != nil {
			log.Println("创建视频失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", video.Author_id).Update("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
			log.Println("作品数更新失败")
			return err
		}
		return nil
	})
}

// 插入关注数据
func Tran_InsertFollow(foDate *Dfollow) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Dfollow{}).Create(foDate).Error; err != nil {
			log.Println("关注数据插入失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", foDate.User_id).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			log.Println("作品数更新失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", foDate.Follow_id).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			log.Println("作品数更新失败")
			return err
		}
		return nil
	})
}

// 关注erase操作  根据关注用户和被关注用户id进行访问
func Tran_EraseFollow(user_id int64, follow_id int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND follow_id = ?", user_id, follow_id).Delete(&Dfollow{}).Error; err != nil {
			log.Println("关注数据插入失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", user_id).Update("follow_count", gorm.Expr("follow_count + ?", -1)).Error; err != nil {
			log.Println("作品数更新失败")
			return err
		}

		if err := tx.Model(&Duser{}).Where("id = ?", follow_id).Update("follower_count", gorm.Expr("follower_count + ?", -1)).Error; err != nil {
			log.Println("作品数更新失败")
			return err
		}
		return nil
	})
}
