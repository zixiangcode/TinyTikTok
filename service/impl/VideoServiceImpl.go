package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
)

type VideoServiceImpl struct {
}

func (videoServiceImpl VideoServiceImpl) GetVideoListByUserID(userID int64) ([]models.Video, error) {
	videoList, err := dao.GetVideoListByUserID(userID)
	if err != nil {
		return []models.Video{}, err
	}
	return videoList, err
}
