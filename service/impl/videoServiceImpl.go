package impl

import (
	"TinyTikTok/config"
	"TinyTikTok/dao"
	"TinyTikTok/db"
	"TinyTikTok/models"
	"TinyTikTok/service"
	"TinyTikTok/utils"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/savsgio/gotils/uuid"

	"mime/multipart"
)

//const AccessKeyId = ""
//const AccessKeySecret = ""
//const  Endpoint = "https://oss-cn-hangzhou.aliyuncs.com" // OSS的访问域名   杭州
//
//const BucketName = "web-tlias-amireux"
//
////const url = "https://"+config.BucketName +".oss-cn-hangzhou.aliyuncs.com/"
// url = "https://"+config.Config.BucketName +".oss-cn-hangzhou.aliyuncs.com/"

type VideoServiceImpl struct {
	service.VideoService
}

// Publish
// 将传入的视频流保存在服务器中，并将链接存储在mysql表中
func (videoService *VideoServiceImpl) Publish(data *multipart.FileHeader, userId int64, title string) error {

	fmt.Println("方法data.Open() 成功")
	//fmt.Println("文件名字叫做",file)
	//生成一个uuid作为视频的名字
	videoName := uuid.V4()
	fmt.Println("视频的名字叫做", videoName)

	//上传视频
	err := utils.UploadToServer(data, videoName)
	if err != nil {
		return err
	}

	//videoName是视频名字		userid是用户id  title是文章的标题
	err = SaveVideo(videoName, userId, title)
	if err != nil {
		log.Println("新的数据插入失败")
		return err
	} else {
		log.Println("新的数据插入成功")
	}
	return nil

}

// SaveVideo 将视频的数据插入到数据库里面
func SaveVideo(name string, userId int64, title string) error {

	var count int64
	//获得人数
	result := db.GetMysqlDB().Model(&models.Video{}).Count(&count).Error
	if result != nil {
		log.Println(result)
	}

	count++
	log.Println("count=", count)

	s := "insert into videos(id ,author_id,play_url,cover_url,favorite_count,comment_count,is_deleted,title,create_time)  values (?,?,?,?,?,?,?,?,?) "
	r, err := db.GetMysql().Exec(s, count, userId, config.VideoConfig.Url+name+".mp4", config.VideoConfig.Url+name+".jpg", 0, 0, false, title, time.Now())

	if err != nil {
		// fmt.Println("插入出现问题")
		log.Printf("插入出现问题err: %v\n", err)
		return err
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}

	s = "update user set work_count=work_count+1 where id=?"
	r, err = db.GetMysql().Exec(s, userId)
	if err != nil {
		fmt.Println("更新失败,err", err)
		return err
	}
	return nil

}

func (videoService *VideoServiceImpl) ShowVideoList(userId int64) ([]models.Video, error) {

	videos, err := QueryVideosById(userId)
	if err != nil {
		log.Println("查询失败,err=", err)
		return nil, err
	}
	return videos, nil

}

func (videoService VideoServiceImpl) GetVideoListByUserID(userID int64) ([]models.Video, error) {
	videoList, err := dao.GetVideoListByUserID(userID)
	if err != nil {
		return []models.Video{}, err
	}
	return videoList, err
}

// QueryVideosById 根据用户id查询视频
func QueryVideosById(user_id int64) ([]models.Video, error) {
	var videos []models.Video

	log.Println("user_id=", user_id)
	s := "select * from videos join user on videos.author_id=user.id where user.id=? and tinytiktok.user.is_deleted=false"
	r, err := db.GetMysql().Query(s, user_id)
	fmt.Println("sql语句是", s)
	var v models.Video
	defer func(r *sql.Rows) {
		err := r.Close()
		if err != nil {

		}
	}(r)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	} else {
		for r.Next() {

			err := r.Scan(&v.Id, &v.CreateTime, &v.IsDeleted, &v.AuthorID, &v.CoverURL, &v.PlayURL, &v.Title, &v.CommentCount, &v.FavoriteCount,
				&v.Author.Id, &v.Author.Name, &v.Author.FollowCount, &v.Author.FollowerCount, &v.Author.IsFollow, &v.Author.Avatar, &v.Author.BackgroundImage,
				&v.Author.Signature, &v.Author.TotalFavorited, &v.Author.WorkCount, &v.Author.FavoriteCount, &v.Author.IsDeleted, &v.Author.CreateTime, &v.Author.Password)
			if err != nil {
				return nil, err
			}
			//log.Println(v)
			videos = append(videos, v)
		}
	}
	defer func(r *sql.Rows) {
		err := r.Close()
		if err != nil {

		}
	}(r)
	return videos, nil

}
