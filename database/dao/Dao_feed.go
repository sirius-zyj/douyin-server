package dao

import (
	// "gorm.io/gorm"

	// "errors"

	"log"
	"time"

	"gorm.io/gorm"
)

func (Dvideo) TableName() string {
	return "videos"
}

// GetVideoById 根据视频的ID获取视频信息
func GetVideoById(id int64) (Dvideo, error) {
	var resp Dvideo
	err := db.Model(&Dvideo{}).Where("id = ?", id).Find(&resp).Error
	if err != nil {
		log.Println(err.Error())
		return resp, err
	}
	return resp, nil
}

// 根据一组视频ID获取视频切片
func GetVideosByIds(ids []int64) ([]Dvideo, error) {
	var resq []Dvideo
	err := db.Model(&Dvideo{}).Where("id in ?", ids).Find(&resq).Error
	if err != nil {
		log.Println("查询失败")
		return resq, err
	}
	return resq, nil
}

// 根据作者的id查询他发布的所有视频
func GetVideoByAuthorId(id int64) ([]Dvideo, error) {
	var resp []Dvideo
	err := db.Model(&Dvideo{}).Where("author_id = ?", id).Find(&resp).Error
	if err != nil {
		log.Println(err.Error())
		return resp, err
	}
	return resp, nil
}

// 根据一个时间获取前面的视频
func GetVideoByTime(time time.Time) ([]Dvideo, error) {
	var resp []Dvideo
	year := time.Year()
	if year < 1 || year > 9999 {
		// 处理年份超出范围的情况
		log.Println("Invalid year:", year)
	} else {
		// 处理年份在有效范围内的情况
		log.Println("Valid year:", year)
	}
	err := db.Model(&Dvideo{}).Where("upload_time <= ?", time).Order("upload_time desc").Limit(30).Find(&resp).Error
	if err != nil {
		log.Println(err.Error())
		return resp, err
	}
	return resp, nil
}

func UpdateFeed(where string, where_count int64, what string, what_count int64) error {
	err := db.Model(&Dvideo{}).Where(where+" = ?", where_count).Update(what, gorm.Expr(what+" + ?", what_count)).Error
	if err != nil {
		log.Println("更新失败")
		return err
	}
	return nil
}

// func UpdateVideoFavoritedCount(video_id int64, count int64) error {
// 	err := db.Model(&Dvideo{}).Where("id = ?", video_id).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", count)).Error
// 	if err != nil {
// 		log.Println("更新失败")
// 		return err
// 	}
// 	return nil
// }

// func UpdateCommentCount(video_id int64, count int64) error {
// 	err := db.Model(&Dvideo{}).Where("id = ?", video_id).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", count)).Error
// 	if err != nil {
// 		log.Println("更新失败")
// 		return err
// 	}
// 	return nil
// }

//通过ftp服务将视频传入服务器

//将图片传给FTP服务器

//保存视频的记录

// func UploadVideo(video *[]byte) (playUrl, coverUrl string, err error) {
// 	videoName := uuid.NewV4().String() + ".mp4"
// 	imageName := uuid.NewV4().String() + ".jpeg"

// 	err = os.WriteFile(videoName, *video, 0o666)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	imageData, _ := GetSnapshot(videoName, 1)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	err = VideoBucket.PutObject(videoName, bytes.NewReader(*video))
// 	if err != nil {
// 		return "", "", err
// 	}

// 	err = ImageBucket.PutObject(imageName, imageData)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	playUrl = VideoBucketLinkPrefix + videoName
// 	coverUrl = ImageBucketLinkPrefix + imageName
// 	return playUrl, coverUrl, nil
// }

// func GetSnapshot(videoPath string, frameNum int) (cover io.Reader, err error) {
// 	buf := bytes.NewBuffer(nil)
// 	err = ffmpeg.Input(videoPath).
// 		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
// 		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
// 		WithOutput(buf, os.Stdout).
// 		Run()

// 	if err != nil {
// 		log.Fatal("Extract Frame Failed", err)
// 		return nil, err
// 	}
// 	err = os.RemoveAll(videoPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bytes.NewReader(buf.Bytes()), nil
// }
