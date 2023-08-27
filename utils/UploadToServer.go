package utils

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func UploadToServer( data *multipart.FileHeader,videoName string) error {

	//将视频流上传到视频服务器，保存视频链接
	file, err := data.Open()
	if err != nil {
		fmt.Printf("方法data.Open() 失败%v\n", err)
		return err
	}
	defer file.Close()

	//-------------------------------
	//创建OSS客户端

	//这里记得把账号密码给填上去
	//client, err := oss.New(config.VideoConfig.Endpoint, config.VideoConfig.AccessKeyId, config.VideoConfig.AccessKeySecret)
	////client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return err
	//}
	//
	//
	//// 获取存储空间
	//bucket, err := client.Bucket(config.VideoConfig.BucketName)
	////bucket, err := client.Bucket(BucketName)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return err
	//}
	//
	//// 上传视频流到OSS
	//err = bucket.PutObject(videoName+".mp4", file)//文件名   后缀添加.mp4
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return err
	//}
	//fmt.Println("上传视频到oos成功")


	//获取第一帧的图片，放到image里面
	err, path := GetImage(videoName)
	if err != nil {
		log.Println("getImage Error:", err)
		return err
	}
	fmt.Println("保存文件的path=",path)

	//打开本地文件
	file, err = os.Open(path)
	if err != nil {
		return err
	}

	//上传文件到 OSS
	//err = bucket.PutObject(videoName+".jpg", file)
	//if err != nil {
	//	return err
	//}



	var wg	sync.WaitGroup
	wg.Add(2)			//开两个协程
	//删除在image里面的文件流
	go  deleteFile(&wg,path)
	go deleteFile(&wg,path + ".jpg")
	wg.Wait()

	if err!=nil{
		return err
	}
	return nil
}

func deleteFile(wg *sync.WaitGroup,path string)error{
	defer wg.Done()
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
func printDir(path string){
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Directory: %s\n", file.Name())
		} else {
			fmt.Printf("File: %s\n", file.Name())
		}
	}
}

//获取图片的第一帧			返回错误和路径
func GetImage(Filename string) (error,string){
	//videoURL := config.VideoConfig.Url+Filename+".mp4"
	videoURL := "https://www.w3schools.com/html/movie.mp4"

	outputPath := "public/"+Filename//图片的路径
	fmt.Println("链接是",videoURL)
	//fmt.Println("--------------")
	//printDir("public/")
	//fmt.Println("--------------")

	//打印当前路径
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("当前工作路径：", dir)
	//注意下面需要配置一下ffmpeg的路径   		查看版本命令		ffmpeg -version
	// 构建FFmpeg命令
	cmd := exec.Command("ffmpeg", "-i", videoURL, "-vframes", "1", "-f", "image2", outputPath)

	// 执行FFmpeg命令
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("视频链接地址错误Error executing FFmpeg command:", err)
		return err, ""
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
	outputFile, err := os.Create("public/"+Filename+".jpg")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return nil, ""
	}
	defer outputFile.Close()

	jpeg.Encode(outputFile, img, nil)

	fmt.Println("Screenshot saved to"+outputPath+" output.jpg")
	return nil, "public/"+Filename
}
