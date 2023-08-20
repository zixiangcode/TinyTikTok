package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
	"TinyTikTok/service"
	"fmt"

	"log"

	"github.com/savsgio/gotils/uuid"

	"mime/multipart"
)


type VideoServiceImpl struct {
	service.VideoService
}


// Publish
// 将传入的视频流保存在服务器中，并将链接存储在mysql表中
func (videoService *VideoServiceImpl) Publish(data *multipart.FileHeader, userId int64, title string) error {
	//将视频流上传到视频服务器，保存视频链接
	file, err := data.Open()
	if err != nil {
		fmt.Printf("方法data.Open() 失败%v\n", err)
		return err
	}
	fmt.Println("方法data.Open() 成功")
	fmt.Println("文件名字叫做",file)
	//生成一个uuid作为视频的名字
	videoName := uuid.V4()
	fmt.Println("视频的名字叫做",videoName)
	err = dao.SaveVideo(data,videoName)
	if err != nil {
		fmt.Printf("方法dao.SaveVideo(data)%v\n", err)
		return err
	}
	fmt.Println("方法dao.SaveVideo(data) 成功")
	defer file.Close()




	err = dao.InsertVideo(videoName,userId,title)
	if err!=nil{
		log.Println("新的数据插入失败")
		return err
	}else{
		log.Println("新的数据插入成功")
	}
	return nil

}



//查询数据列表
func (videoService *VideoServiceImpl) ShowVideoList(id int64)([]models.Video,error) {

	videos, err := dao.QueryVideoById(id)

	if err!=nil{
		log.Println("查询数据失败了，error=",err)
		return nil,err
	}else{
		//fmt.Println(videos)
		return videos,nil
	}
}

//查询用户是否存在
func (videoService *VideoServiceImpl)IsExist(userId int64)(bool,error){
	exist, err := dao.QueryUserIsExist(userId)
	if err!=nil{
		return  false,err
	}else{
		return exist,nil
	}
}


