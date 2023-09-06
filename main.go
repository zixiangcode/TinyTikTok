package main

import (
	"TinyTikTok/config"
	"TinyTikTok/db"
	"TinyTikTok/router"
	"TinyTikTok/service"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ReadConfig("config/configuration.json") // 先读取配置文件
	config.ReadVideoServerConfig("config/videoserverconfig.json")
	db.CreateGORMDB() // 创建 GORM 连接 MySql

	go service.RunMessageServer()

	r := gin.Default()

	router.InitRouter(r)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
