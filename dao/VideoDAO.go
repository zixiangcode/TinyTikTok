package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
)

func GetVideoListByUserID(UserID int64) ([]models.Video, error) {
	var videoList []models.Video
	result := db.GetMysqlDB().
		Where("author_id = ?", UserID).
		Order("create_time DESC").
		Find(&videoList)
	return videoList, result.Error
}

func GetUserByVideoID(videoID int64) (models.User, error) {
	var user models.User
	result := db.GetMysqlDB().
		Where("id = ?", videoID).
		Order("create_time DESC").
		Find(&user)
	return user, result.Error
}
