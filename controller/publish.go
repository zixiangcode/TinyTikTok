package controller

import (
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

func GetVideoService() impl.VideoServiceImpl { //创建一个视频流接口
	var VideoService impl.VideoServiceImpl
	return VideoService
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	data, err := c.FormFile("data") //获取上传的数据
	if err != nil {
		log.Printf("上传数据出现问题\n")

		c.JSON(http.StatusOK, models.Response{ //返回错误信息
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	} else {
		log.Println("上传没问题")
	}
	//获取视频名字
	videoName := filepath.Base(data.Filename)
	log.Println("视频文件的名字为", videoName)

	token := c.PostForm("token")

	//userId, err := strconv.ParseInt(token, 10, 64)//用户账号
	if err != nil {
		//log.Println("转化出问题了")
		c.JSON(http.StatusOK, models.Response{ //返回错误信息
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	fmt.Println("用户账号是", token)
	title := c.PostForm("title")
	fmt.Println("标题是" + title)

	// 将截取的token转化为userID

	userClaim, err := utils.ParseToken(token)

	userID := userClaim.JWTCommonEntity.Id

	fmt.Println("userID=", userID)

	//获取接口
	videoService := GetVideoService()

	//上传文件
	err = videoService.Publish(data, userID, title)
	if err != nil {
		fmt.Printf("videoService.Publish(data, userId) 失败：%v\n", err)
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//log.Println("videoService.Publish(data, userId) 成功")
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	//获取user_id
	query := c.Query("user_id")
	username, err := strconv.ParseInt(query, 10, 64)

	serviceImpl := GetVideoService() //创建接口
	list, err := serviceImpl.ShowVideoList(username)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: models.Response{
				StatusMsg:  "查询失败",
				StatusCode: 1,
			},
			VideoList: nil,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: models.Response{
				StatusMsg:  "查询成功",
				StatusCode: 0,
			},
			VideoList: list,
		})
	}
}
