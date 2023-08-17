package dao

import (
  "errors"
  "log"
  "gorm.io/gorm"
  "database/sql"
)
func(Dfavorite) TableName() string {
  return "favorite";
}

//点赞更新操作  根据点赞用户和视频id进行访问
func ActionFavorite(user_id int64 , video_id int64 , action bool) error {
  err := db.Model(&Dfavorite{}).Where("user_id = ? AND video_id = ?" , user_id , video_id).Update("cancel" , action).Error

  if err != nil {
    log.Println("点赞更新错误")
    return errors.New("点赞更新错误")
  }
  return nil
}

//插入点赞数据
func InsertFavorite(faDate Dfavorite) error{
  err := db.Model(&Dfavorite{}).Create(&faDate).Error
  if err != nil {
    log.Println("点赞数据插入失败")
    return errors.New("点赞数据插入失败")
  }
  return nil
}

//查询点赞信息
func GetFavoriteData(user_id int64 , video_id int64) (Dfavorite , error) {
  var favoriteData Dfavorite
  err := db.Model(&Dfavorite{}).Where("user_id = ? AND video_id = ?" , user_id , video_id).Find(&favoriteData).Error
  if err != nil {
    log.Println("查询错误")
    favoriteData.Id = -1
    return favoriteData , errors.New("查询出错")
  }
  return favoriteData , nil
}


//查询某用户的点赞视频ID列表
func GetFavoriteList(user_id int64) ([]int64 , error) {
  var res []sql.NullInt64
  var resq []int64
  err := db.Model(&Dfavorite{}).Where("user_id = ?" , user_id).Pluck("video_id" , &res).Error
  for _ , tmp := range(res) {
    if tmp.Valid {
      resq = append(resq , tmp.Int64)
    }
  }
  if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      // 查询为空
      log.Println("查询为空")
      return resq , nil
    } else {
      // 其他错误
      log.Println("喜欢视频数据查询发生错误")
      return resq , errors.New("喜欢视频数据查询发生错误")
    }
  }
  return resq , nil
}