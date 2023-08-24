package service

import "TinyTikTok/models"

type FeedService interface {
	GetFeedLatestTime(lastTime string) (int64, error)                                   //获取时间戳
	GetFeedByLatestTime(timestamp int64) ([]models.FeedResponseVideoInfo, int64, error) //获取视频流
}
