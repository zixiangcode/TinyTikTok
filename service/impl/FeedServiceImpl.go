package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
	"log"
	"strconv"
	"time"
)

type FeedServiceImpl struct {
}

func (feedServiceImpl FeedServiceImpl) GetFeedLatestTime(lastTime string) (int64, error) {

	if lastTime != "" && lastTime != "0" {
		latestTime, err := strconv.ParseInt(lastTime, 10, 64)
		if err != nil {
			log.Printf("方法GetFeedLatestTime失败：%v", err)
			return 1, err
		}

		//app打开第一次的时间戳为毫秒级，需要改为秒级
		if len(lastTime) == 13 {
			latestTime /= 1000
		}

		return latestTime, err
	}
	return time.Now().Unix(), nil
}

func (feedServiceImpl FeedServiceImpl) GetFeedByLatestTime(timestamp int64) ([]models.FeedResponseVideoInfo, int64, error) {

	//将时间戳格式数据改为time.Time格式
	latestTime := time.Unix(timestamp, 0)

	//根据时间戳从数据库中查询video
	videos, err := dao.GetFeedVideosInfo(latestTime)
	if err != nil {
		log.Printf("方法GetFeedByLatestTime失败：%v", err)
		return []models.FeedResponseVideoInfo{}, 0, err
	}
	var feedResponseVideoInfos = make([]models.FeedResponseVideoInfo, len(videos))
	for k, video := range videos {

		//从videos中读取userID查询uesr信息，并拼接到response中
		user, err := UserServiceImpl{}.GetUserById(video.AuthorID)
		if err != nil {
			log.Printf("方法 GetFeedByLatestTime 失败 %v", err)
			return []models.FeedResponseVideoInfo{}, 0, err
		}
		//查询每个video的评论总数
		comments, err := dao.GetCommentsByVideoID(video.Id)
		if err != nil {
			log.Printf("方法 GetFeedByLatestTime 失败 %v", err)
			return []models.FeedResponseVideoInfo{}, 0, err
		}
		commentCount := int64(len(comments))

		//todo 查询每个video的点赞总数
		favoriteCount := int64(1)

		//todo 判断当前用户是否关注本条视频
		isFavorite := false

		var feedResponseVideoInfo = models.FeedResponseVideoInfo{
			Author: models.FeedUserInfo{
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				FavoriteCount:   user.FavoriteCount,
				FollowCount:     user.FollowCount,
				FollowerCount:   user.FollowerCount,
				ID:              user.Id,
				IsFollow:        user.IsFollow,
				Name:            user.Name,
				Signature:       user.Signature,
				TotalFavorited:  user.TotalFavorited,
				WorkCount:       user.WorkCount,
			},
			CommentCount:  commentCount,
			CoverURL:      video.CoverURL,
			FavoriteCount: favoriteCount,
			ID:            video.Id,
			IsFavorite:    isFavorite,
			PlayURL:       video.PlayURL,
			Title:         video.Title,
		}
		feedResponseVideoInfos[k] = feedResponseVideoInfo
	}
	//获得下次查询视频的时间戳
	nextTime := videos[len(videos)-1].CreateTime.Unix()
	return feedResponseVideoInfos, nextTime, err
}
