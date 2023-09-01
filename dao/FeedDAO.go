package dao

import (
	"TinyTikTok/config"
	"TinyTikTok/db"
	"TinyTikTok/models"
	"time"
)

func GetFeedVideosInfo(latestTime time.Time) ([]models.Video, error) {
	var videos []models.Video
	result := db.GetMysqlDB().
		Where("create_time < ? AND is_deleted <> 1", latestTime).
		Order("create_time desc").
		Limit(config.VideoCount).
		Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}
