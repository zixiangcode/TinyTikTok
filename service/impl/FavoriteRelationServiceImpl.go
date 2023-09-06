package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
	"errors"
)

type FavoriteRelationServiceImpl struct {
}

func (favoriteRelationServiceImpl FavoriteRelationServiceImpl) AddfavoriteRelation(favoriteRelation models.FavoriteRelation) error {
	//判断是否已点赞，如果已点赞就不修改user中的favorite_count,video中的favorite_count
	if favoriteRelationServiceImpl.IsFavorite(favoriteRelation.UserID, favoriteRelation.VideoID) {
		return errors.New("已点赞")
	}

	err := dao.AddfavoriteRelation(favoriteRelation)
	return err
}

func (favoriteRelationServiceImpl FavoriteRelationServiceImpl) DeletefavoriteRelation(favoriteRelation models.FavoriteRelation) error {
	//判断是否已取消点赞，如果已取消点赞或者未点赞就不修改user中的favorite_count,video中的favorite_count
	if !favoriteRelationServiceImpl.IsFavorite(favoriteRelation.UserID, favoriteRelation.VideoID) {
		return errors.New("未点赞或已取消点赞")
	}

	err := dao.DeletefavoriteRelation(favoriteRelation)
	return err
}

func (favoriteRelationServiceImpl FavoriteRelationServiceImpl) IsFavorite(userID int64, videoID int64) bool {
	favoriteRelation, err := dao.IsfavoriteRelation(userID, videoID)
	if err != nil || (favoriteRelation == models.FavoriteRelation{}) {
		return false
	}
	return true
}

func (favoriteRelationServiceImpl FavoriteRelationServiceImpl) GetFavoriteRelationListByUserID(userID int64) ([]models.FavoriteVideoInfo, error) {
	//1.根据userID查询所有的点赞视频列表
	favoritevideoList, err := VideoServiceImpl{}.GetVideoListByUserID(userID)
	if err != nil {
		return []models.FavoriteVideoInfo{}, err
	}
	//2.查询每个视频的作者信息
	var FavoriteVideoInfoList []models.FavoriteVideoInfo
	for _, video := range favoritevideoList {
		user, err := UserServiceImpl{}.GetUserById(video.AuthorID)
		if err != nil {
			return []models.FavoriteVideoInfo{}, err
		}

		//获取评论视频作者ID用于查询是否关注
		videoUser, err := VideoServiceImpl{}.GetUserByVideoID(video.AuthorID)
		if err != nil {
			return []models.FavoriteVideoInfo{}, err
		}

		//拼接结果
		favoriteVideoInfo := models.FavoriteVideoInfo{
			ID: video.Id,
			Author: models.FavoriteUserInfo{
				ID:              user.Id,
				Name:            user.Name,
				FollowCount:     user.FollowCount,
				FollowerCount:   user.FollowerCount,
				IsFollow:        RelationServiceImpl{}.IsFollow(user.Id, videoUser.Id),
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				TotalFavorited:  user.TotalFavorited,
				WorkCount:       user.WorkCount,
				FavoriteCount:   user.FavoriteCount,
			},
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    FavoriteRelationServiceImpl{}.IsFavorite(userID, video.Id),
			Title:         video.Title,
		}
		FavoriteVideoInfoList = append(FavoriteVideoInfoList, favoriteVideoInfo)
	}
	return FavoriteVideoInfoList, nil
}
