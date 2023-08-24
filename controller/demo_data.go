package controller

import (
	"TinyTikTok/models"
	"time"
)

var DemoVideos = []models.Video{
	{
		CommonEntity: models.CommonEntity{Id: 1},
		//Author:        DemoUser,
		PlayURL:       "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		//CommentCount:  0,
		//IsFavorite:    false,
	},
}

var DemoComments = []models.Comment{
	{
		CommonEntity: models.CommonEntity{Id: 1, CreateTime: time.Now()},
		UserID:       1,
		Content:      "Test Comment",
	},
}

var DemoUser = models.User{
	CommonEntity:  models.CommonEntity{Id: 1},
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
