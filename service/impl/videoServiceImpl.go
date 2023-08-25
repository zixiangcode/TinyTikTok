package impl

import (
	"TinyTikTok/config"
	"TinyTikTok/db"
	"TinyTikTok/models"
	"TinyTikTok/service"
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"log"

	"github.com/savsgio/gotils/uuid"

	"mime/multipart"
)

const AccessKeyId = ""
const AccessKeySecret = ""
const  Endpoint = "https://oss-cn-hangzhou.aliyuncs.com" // OSS的访问域名   杭州

const BucketName = "web-tlias-amireux"

//const url = "https://"+config.BucketName +".oss-cn-hangzhou.aliyuncs.com/"
const url = "https://"+BucketName +".oss-cn-hangzhou.aliyuncs.com/"


type VideoServiceImpl struct {
	service.VideoService
}
const FilePath = "http://localhost:8080/"

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
	//fmt.Println("文件名字叫做",file)
	//生成一个uuid作为视频的名字
	videoName := uuid.V4()
	fmt.Println("视频的名字叫做",videoName)
	err = SaveVideo(data,videoName)
	if err != nil {
		fmt.Printf("方法dao.SaveVideo(data)%v\n", err)
		return err
	}
	fmt.Println("方法dao.SaveVideo(data) 成功")
	defer file.Close()

	err = InsertVideo(videoName,userId,title)
	if err!=nil{
		log.Println("新的数据插入失败")
		return err
	}else{
		log.Println("新的数据插入成功")
	}
	return nil

}


//上传视频文件同时保存第一帧
func SaveVideo(fileHeader *multipart.FileHeader, Filename string) error {
	// 打开视频流文件
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("打开文件失败")
		return err
	}
	defer file.Close()
	//-------------------------------
	// 创建OSS客户端
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	//client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}


	// 获取存储空间
	bucket, err := client.Bucket(config.BucketName)
	//bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// 上传视频流到OSS
	err = bucket.PutObject(Filename+".mp4", file)//文件名   后缀添加.mp4
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("上传视频到oos成功")



	//获取第一帧的图片，放到image里面
	err, path := GetImage(Filename)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// 打开本地文件
	file, err = os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 上传文件到 OSS
	err = bucket.PutObject(Filename+".jpg", file)
	if err != nil {
		return err
	}




	//删除在image里面的文件流
	deleteFile(path)
	deleteFile(path+".jpg")

	return nil

}
func deleteFile(path string)error{
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问文件或目录时出错：%v\n", err)
			return err
		}

		if !info.IsDir(){
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("删除图片时出错：%v\n", err)
				return err
			}

			fmt.Printf("已成功删除图片：%s\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时出错：%v\n", err)
		return err
	}
	return nil
}


//获取图片的第一帧			返回错误和路径
func GetImage(Filename string) (error,string){
	videoURL := url+Filename+".mp4"
	outputPath := "image/"+Filename//图片的路径
	fmt.Println("链接是",videoURL)

	//注意下面需要配置一下ffmpeg的路径   		查看版本命令		ffmpeg -version
	// 构建FFmpeg命令
	cmd := exec.Command("ffmpeg", "-i", videoURL, "-vframes", "1", "-f", "image2", outputPath)

	// 执行FFmpeg命令
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing FFmpeg command:", err)
		return nil, ""
	}


	// 读取输出文件
	file, err := os.Open(outputPath)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return nil, ""
	}
	defer file.Close()

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading output file:", err)
		return nil, ""
	}


	// 将文件内容转化为图片流
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil, ""
	}


	// 调整图片尺寸
	img = resize.Resize(1280, 850, img, resize.Lanczos3)

	// 创建图片
	outputFile, err := os.Create("image/"+Filename+".jpg")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return nil, ""
	}
	defer outputFile.Close()

	jpeg.Encode(outputFile, img, nil)

	fmt.Println("Screenshot saved to"+outputPath+" output.jpg")
	return nil, "image/"+Filename
}


//--------------------------

//将视频的数据插入到数据库里面
func InsertVideo(name string, userId int64, title string) error {


	currentTime := time.Now()

	// 获取当前时间的Unix时间戳
	unixTime := currentTime.Unix()

	// 将Unix时间戳转换为整数
	intValue := int(unixTime)

	//获得人数
	//result :=db.GetMysqlDB().Model(&models.Video{}).Count(&count).Error
	//if result!=nil{
	//	log.Println(result)
	//}
	var count int64
	err := db.GetMysql().QueryRow("SELECT COUNT(*) FROM videos").Scan(&count)
	if err != nil {
		log.Println(err)
	}
	count++
	log.Println("count=",count)
	//fmt.Println(user)
	//video := models.Video{
	//	//CommonEntity: entity,
	//	Author: user,
	//	//ID: 6,
	//	ID:count,
	//	CreateDate: int64(intValue),
	//	UserId:userId,
	//	CommentCount: 0,
	//	CoverURL:url+name+".jpg",
	//	FavoriteCount: 0,
	//	IsFavorite: false,
	//	PlayURL: url+name+".mp4",
	//	Title: title,
	//}
	//fmt.Println(video)
	//log.Println("上面的走完了")
	//result = db.GetMysqlDB().Create(&video).Error
	//if result.Error!=nil{
	//	log.Println("插入出错了",result)
	//	return result
	//}
	//上面这个版本不知道为啥会导致内存泄露
	s := "insert into videos(id ,user_id,play_url,cover_url,favorite_count,comment_count,is_favorite,title,create_date) values (?,?,?,?,?,?,?,?,?) "
	r, err := db.GetMysql().Exec(s, count,userId, url+name+".mp4",url+name+".jpg",0,0,false,title,int64(intValue))

	if err != nil {
		// fmt.Println("插入出现问题")
		fmt.Printf("插入出现问题err: %v\n", err)
		return err
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}

	return nil

}


func(videoService *VideoServiceImpl) ShowVideoList(userId int64) ([]models.Video, error){

	videos, err := QueryVideosById(userId)
	if err!=nil {
		log.Println("查询失败,err=", err)
		return nil, err
	}
	return videos,nil

}

func QueryVideosById(user_id int64)  ([]models.Video,error){		//根据用户名查询视频
	var videos []models.Video
	//var err error
	//go func() {
	//	err = db.GetMysqlDB().Preload("User", "id = tinytiktok.videos.user_id").Where("videos.user_id=?", user_id).Find(&videos).Error
	//}()
	//
	//if err!=nil{
	//	log.Println("查询语句出现了问题，err",err)
	//	return nil,err
	//}
	//log.Println("查询没啥问题")
	//log.Println(videos)
	var tempint int

 	s := "select * from videos join user on tinytiktok.videos.user_id=user.id where user.id=?	 and tinytiktok.user.is_deleted=false"
	r, err := db.GetMysql().Query(s,user_id)
	fmt.Println("sql语句是",s)
	var v models.Video
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil,err
	} else {
		for r.Next() {

			r.Scan(&v.Id,&v.Author.Id,&v.PlayURL,&v.CoverURL,&v.FavoriteCount,&v.CommentCount,&v.IsFavorite,&v.Title,&v.CreateDate,&tempint,
				&v.Author.Name,&v.Author.FollowCount,&v.Author.FollowerCount,&v.Author.IsFollow,&v.Author.Avatar,&v.Author.BackgroundImage,
				&v.Author.Signature,&v.Author.TotalFavorited,&v.Author.WorkCount,&v.Author.FavoriteCount,&tempint,&tempint, &tempint)
			log.Println(v)
			videos=append(videos, v)
		}
	}
	return videos,nil
	defer r.Close()
	return videos,nil

}
