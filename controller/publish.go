package controller

import (
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

func createVideoServiceImpl() impl.VideoServiceImpl { //创建一个视频流接口
	var VideoService impl.VideoServiceImpl
	return VideoService
}


// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	data, err := c.FormFile("data")//获取上传的数据
	if err!=nil{
		fmt.Printf("上传数据出现问题\n")
<<<<<<< Updated upstream
		c.JSON(http.StatusOK, Response{//返回错误信息
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}else{
		fmt.Println("上传没问题")
	}
	//获取视频名字
	videoName := filepath.Base(data.Filename)
	fmt.Println("视频文件的名字为",videoName)


	token:=c.PostForm("token")

	fmt.Printf("id=%v  类型是%v\n",token,token)
	//userId, err := strconv.ParseInt(token, 10, 64)//用户账号
	if err!=nil{
		fmt.Println("转化出问题了")
	}
	fmt.Println("用户账号是",token)
	title := c.PostForm("title")
	fmt.Println("标题是"+title)

	      //后期这里改成jwt解密
	      var userID int64
		userID=123456789


		fmt.Println("userID=",userID)


	//获取接口
	videoService := createVideoServiceImpl()

	exist, err := videoService.IsExist(userID)//查询用户是否存在
	if err!=nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "未知错误",
		})
		return
	}else if !exist {//没查询到就是0
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "该用户不存在",
		})
=======
		c.JSON(http.StatusOK, models.Response{//返回错误信息
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
>>>>>>> Stashed changes
		return
	}else{
		fmt.Println("上传没问题")
	}
	//获取视频名字
	videoName := filepath.Base(data.Filename)
	fmt.Println("视频文件的名字为",videoName)

<<<<<<< Updated upstream
=======

	token:=c.PostForm("token")

	fmt.Printf("id=%v  类型是%v\n",token,token)
	//userId, err := strconv.ParseInt(token, 10, 64)//用户账号
	if err!=nil{
		fmt.Println("转化出问题了")
	}
	fmt.Println("用户账号是",token)
	title := c.PostForm("title")
	fmt.Println("标题是"+title)

	//后期这里改成jwt解密
	var userID int64
	userID=1


	fmt.Println("userID=",userID)


	//获取接口
	videoService := createVideoServiceImpl()



>>>>>>> Stashed changes
	//上传文件
	err = videoService.Publish(data, userID, title)
	if err != nil {
		fmt.Printf("videoService.Publish(data, userId) 失败：%v\n", err)
<<<<<<< Updated upstream
		c.JSON(http.StatusOK, Response{
=======
		c.JSON(http.StatusOK, models.Response{
>>>>>>> Stashed changes
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	fmt.Println("videoService.Publish(data, userId) 成功")

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	videoService := createVideoServiceImpl()//创建流


	q := c.Query("token")
	//var UserId int64
	//UserId=123456789
	//fmt.Println(UserId)

	//后期记得把这里jwt解密
	username, err := strconv.ParseInt(q, 10, 64)
	list, err := videoService.ShowVideoList(username)//展示列表

	if err!=nil{
		c.JSON(http.StatusOK, VideoListResponse{
			Response: models.Response{
				StatusCode: 1,
			},
			VideoList: nil,
		})
	} else{
		c.JSON(http.StatusOK, VideoListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			VideoList: list,
		})
	}

}
