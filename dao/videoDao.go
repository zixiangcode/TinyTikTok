package dao

import (
	"TinyTikTok/config"
	"TinyTikTok/models"
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
)


//阿里云oss
 // 您的OSS存储空间名称
//const objectName = "your-object-name"  // 上传到OSS的对象名称
//const localFile = "your-local-video-file"  // 本地视频文件路径

const AccessKeyId = ""
const AccessKeySecret = ""
const  Endpoint = "https://oss-cn-hangzhou.aliyuncs.com" // OSS的访问域名   杭州

const BucketName = "web-tlias-amireux"

//const url = "https://"+config.BucketName +".oss-cn-hangzhou.aliyuncs.com/"
const url = "https://"+BucketName +".oss-cn-hangzhou.aliyuncs.com/"


//上传视频文件同时保存第一帧
func SaveVideo(fileHeader  *multipart.FileHeader,Filename string ) error  {
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


////判断后缀		删除文件用的
//func isImageFile(path string) bool {
//	ext := strings.ToLower(filepath.Ext(path))
//	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
//}

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

//获取上传链接  测试
//----------------
func GetUploadURL(videoname string) (string, error){
	//初始化
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	//client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Println("Error creating OSS client:", err)
		return "",err
	}

	objectName := videoname // 存储对象名称

	bucket, err := client.Bucket(BucketName)
	//bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("Error getting bucket:", err)
		return "",err
	}

	url, err := bucket.SignURL(objectName, oss.HTTPGet, 3600) // 生成一个有效期为1小时的URL
	if err != nil {
		fmt.Println("Error generating signed URL:", err)
		return "",err
	}

	fmt.Println("URL:", url)
	return url,nil
}

func GetUploadFileURL(videoname string) (string, error) {


	fmt.Printf("名字是%v\n",videoname)
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("创建Error:", err)
		return "", err
	}
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret, oss.SetCredentialsProvider(&provider))
	//client, err := oss.New(Endpoint,AccessKeyId,AccessKeySecret, oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("new的Error:", err)
		return "", err
	}

	// 填写Object的完整路径，完整路径中不能包含Bucket名称，例如exampledir/exampleobject.txt。
	objectName := config.BucketName+"/"+videoname
	//objectName := BucketName+"/"+videoname

	bucket, err := client.Bucket(BucketName)
	//bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// 填写Object的VersionId。
	// 生成签名URL，并指定签名URL的有效时间为60秒。
	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, 60, oss.VersionId("CAEQEhiBgIDmgPf8mxgiIDA1YjZlNDIxY2ZmMzQ1MmU5MTM1Y2M4Yzk4******"))
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	fmt.Printf("Sign Url:%s\n", signedURL)

	return signedURL, nil
}
//--------------------------


//将视频的数据插入到数据库里面
func InsertVideo(name string,userId int64,title string)error {

	s := "insert into video(authorId,playUrl,coverUrl,title) values (?,?,?,?) "
	r, err := config.Db.Exec(s, userId, url+name+".mp4",url+name+".jpg",title)

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


//查询作者所有的video数据  	对应list
func QueryVideoById(id int64) ([]models.Video ,error){

	var videos []models.Video
	//先查询user，至于为啥，我写联合查询，后面的接收不了
	//s := "select * from user where id = ?"
	//
	//r, err := config.Db.Query(s,id)
	//fmt.Println("sql语句是",s,id)
	//var u models.User
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return nil,err
	//} else {
	//	for r.Next() {
	//		log.Println("上面的走过")
	//		r.Scan(&u.Id,&u.Name,&u.FollowCount,&u.FollowerCount,&u.IsFollow,&u.Avatar,&u.Background_image,
	//			&u.Signature,&u.Total_favorited,&u.Work_count,&u.Favorite_count)
	//	}
	//}
	//
	//s = "select * from video where authorId=?"
	//r, err = config.Db.Query(s,id)
	//fmt.Println("sql语句是",s,id)
	//
	//var v models.Video
	//
	//
	////var count int
	//if err != nil {
	//	log.Println("出错了")
	//	fmt.Printf("err: %v\n", err)
	//	return nil,err
	//} else {
	//	for r.Next() {
	//		//注意一下这里,要是说fallowCount在数据库里里面为0,它好像不会返回给前端			2023-08-18  	下面的最后面三个不会返回
	//		r.Scan(&v.Id, &v.Author.Id, &v.PlayUrl,&v.CoverUrl,&v.FavoriteCount,&v.CommentCount,&v.IsFavorite,&v.Title,
	//			&v.Next_time)
	//		v.Author=u
	//
	//		videos=append(videos, v)
	//	}
	//}

	s := "select * from video join user on authorId=user.id where user.id=?"
	r, err := config.Db.Query(s,id)
	fmt.Println("sql语句是",s)
	var v models.Video
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil,err
	} else {
		for r.Next() {
			r.Scan(&v.Id, &v.Author.Id, &v.PlayUrl,&v.CoverUrl,&v.FavoriteCount,&v.CommentCount,&v.IsFavorite,
				&v.Title,&v.Next_time,&v.Author.Id,&v.Author.Name,&v.Author.FollowCount,&v.Author.FollowerCount,&v.Author.IsFollow,
				&v.Author.Avatar,&v.Author.Background_image,&v.Author.Signature,&v.Author.Total_favorited,&v.Author.Work_count,&v.Author.Favorite_count)
			videos=append(videos, v)
		}
	}
	return videos,nil
	defer r.Close()
	//log.Println("count",count)
	return videos,nil
}


//查询所有的video数据    刷视频用的   对应feed
func QueryEveryVideo() ([]models.Video ,error){

	var videos []models.Video

	s := "select * from video join user on authorId=user.id"
	r, err := config.Db.Query(s)
	fmt.Println("sql语句是",s)
	var v models.Video
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil,err
	} else {
		for r.Next() {
			r.Scan(&v.Id, &v.Author.Id, &v.PlayUrl,&v.CoverUrl,&v.FavoriteCount,&v.CommentCount,&v.IsFavorite,
				&v.Title,&v.Next_time,&v.Author.Id,&v.Author.Name,&v.Author.FollowCount,&v.Author.FollowerCount,&v.Author.IsFollow,
				&v.Author.Avatar,&v.Author.Background_image,&v.Author.Signature,&v.Author.Total_favorited,&v.Author.Work_count,&v.Author.Favorite_count)
			videos=append(videos, v)
		}
	}
	return videos,nil
}

//查询用户是否存在
func QueryUserIsExist(userid int64)(bool,error){

	query := "SELECT COUNT(*) FROM user WHERE id = ?"
	var count int
	err := config.Db.QueryRow(query, userid).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return false,err
	}
	if count > 0 {
		return true,nil
	} else {
		return false,nil
	}

}
