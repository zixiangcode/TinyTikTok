package dao

import (
	"TinyTikTok/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDB() (err error) {

	dsn := "root:runaway0926@tcp(127.0.0.1:3306)/TinyTikTok?charset=utf8mb4&parseTime=True"
	d, err := sql.Open("mysql", dsn)
	config.Db = d
	if err != nil {
		log.Println("初始化失败",err)
		return err
	}
	err = d.Ping()
	if err != nil {
		log.Println("初始化失败",err)
		return err
	}
	log.Println("初始化成功")
	return nil
}
