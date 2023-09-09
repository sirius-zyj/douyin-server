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
	"gorm.io/plugin/dbresolver"
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

	if db, err = gorm.Open(mysql.New(mysql.Config{DSN: config.MasterDsn}),
		&gorm.Config{Logger: newLogger},
	); err != nil {
		log.Panicln("mysql 连接错误：", err)
	}
	db.AutoMigrate(&Dvideo{}, &Duser{}, &Dfavorite{}, &Dcomments{}, &Dfollow{}, &Dmessage{})

	replicas := []gorm.Dialector{
		mysql.New(mysql.Config{
			DSN: config.SlaveDsn,
		}),
	}
	db.Use(
		dbresolver.Register(dbresolver.Config{
			Sources: []gorm.Dialector{mysql.New(mysql.Config{
				DSN: config.MasterDsn,
			})},
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}).
			SetMaxIdleConns(10).
			SetConnMaxLifetime(time.Hour).
			SetMaxOpenConns(200),
	)

}
