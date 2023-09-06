package dao

import (
	"douyin-server/config"
	"douyin-server/middleware/snowflake"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var SnowFlakeNode *snowflake.Node

func Init() {
	//雪花算法
	var err error
	if SnowFlakeNode, err = snowflake.NewNode(1); err != nil {
		log.Panicln("雪花算法初始化错误：", err)
	}

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
