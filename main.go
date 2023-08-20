package main

import (
	"TinyTikTok/dao"
	"TinyTikTok/router"
	//"TinyTikTok/service"
	"github.com/gin-gonic/gin"
)

type ApifoxModel struct {
	StatusCode int64  `json:"status_code"`// 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"` // 返回状态描述
	Token      string `json:"token"`      // 用户鉴权token
	UserID     int64  `json:"user_id"`    // 用户id
}

func main() {
	//go service.RunMessageServer()

	r := gin.Default()

	router.InitRouter(r)

	//find(r)
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func init(){
	dao.InitDB()
}



