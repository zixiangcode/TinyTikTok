package db

import (
	"TinyTikTok/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var GORM *gorm.DB

func CreateGORMDB() {
	db, err := gorm.Open(mysql.Open(config.Config.Mysql), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(100)                // 连接池最大打开连接数
	sqlDb.SetMaxIdleConns(25)                 // 连接池最大空闲连接数
	sqlDb.SetConnMaxLifetime(1 * time.Minute) // 最大生存时间

	GORM = db // 赋值给全局变量 GORM
}

// GetMysqlDB 需要使用数据库的时候直接创建一个连接 调用此方法即可
func GetMysqlDB() *gorm.DB {
	return GORM
}
