package dao

import (
	"douyin-server/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	//日志系统
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	dsn := config.Dsn
	var err error
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicln("mysql 连接错误：", err)
	}
	db.AutoMigrate(&Dvideo{})
	db.AutoMigrate(&Duser{})
	db.AutoMigrate(&Dfavorite{})
	db.AutoMigrate(&Dcomments{})
	db.AutoMigrate(&Dfollow{})
	db.AutoMigrate(&Dmessage{})
}
