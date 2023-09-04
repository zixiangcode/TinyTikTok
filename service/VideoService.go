package service

import "TinyTikTok/models"

type VideoService interface {
	GetVideoListByUserID(userID int64) ([]models.Video, error)
}
