package service

import (
	"TinyTikTok/models"
	"mime/multipart"
)

type VideoService interface {
	ShowVideoList(userId int64) ([]models.Video, error) //展示自己发了哪些视频

	Publish(data *multipart.FileHeader, userId int64, title string) error //上传视频到服务端

	GetVideoListByUserID(userID int64) ([]models.Video, error)

	GetUserByVideoID(videoID int64) (models.User, error)
}
