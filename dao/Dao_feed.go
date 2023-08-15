package dao

import (
  // "gorm.io/gorm"
  "time"
  "log"
)

func(Dvideo) TableName() string {
  return "videos";
}


//GetVideoById 根据视频的ID获取一系列视频信息
func GetVideoById(id int64) (Dvideo , error) {
  var resp Dvideo
  err := db.Model(&Dvideo{}).Where("id = ?" , id).Find(&resp).Error
  if err != nil {
    log.Println(err.Error())
    return resp , err
  }
  return resp , nil
}

//根据作者的id查询他发布的所有视频
func GetVideoByUserId(id int64) ([]Dvideo , error) {
  var resp []Dvideo
  err := db.Model(&Dvideo{}).Where("author_id = ?" , id).Find(&resp).Error
  if err != nil {
    log.Println(err.Error())
    return resp , err
  }
  return resp , nil
}

//根据一个时间获取前面的视频
func GetVideoByTime(time time.Time) ([]Dvideo , error) {
  var resp []Dvideo
  year := time.Year()
  if year < 1 || year > 9999 {
      // 处理年份超出范围的情况
      log.Println("Invalid year:", year)
  } else {
      // 处理年份在有效范围内的情况
      log.Println("Valid year:", year)
  }
  err := db.Model(&Dvideo{}).Where("upload_time <= ?" , time).Order("upload_time desc").Limit(30).Find(&resp).Error
  if err != nil {
    log.Println(err.Error())
    return resp , err
  }
  return resp , nil
}

//通过ftp服务将视频传入服务器

//将图片传给FTP服务器

//保存视频的记录