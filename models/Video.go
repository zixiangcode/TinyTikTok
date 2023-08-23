package models

import "time"

type Video struct {
<<<<<<< Updated upstream
	//Id            int64  `json:"id,omitempty"`
	CommonEntity
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"isFavorite,omitempty"`
	Title 		string		`json:"title,omitempty"`
	Next_time   int64	`json:"next_time,omitempty"`

}

var DemoVideos = []Video{
	{
		//Id:            8,
		CommonEntity: DemoComEntity,
		Author:        DemoUser,
		//PlayUrl:       "https://web-tlias-amireux.oss-cn-hangzhou.aliyuncs.com/0e148e93-1c67-44ce-aa81-7835e3d62f63.mp4",
		PlayUrl:       "https://web-tlias-amireux.oss-cn-hangzhou.aliyuncs.com/6aff7637-d31a-4f7f-ba7a-b96183886391.mp4",
		//PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://web-tlias-amireux.oss-cn-hangzhou.aliyuncs.com/QQ%E5%9B%BE%E7%89%8720230808172010.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},

}

var DemoComEntity =CommonEntity{
	Id: 123456789,
	CreateTime: time.Now(),
	IsDeleted: 0,

}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
=======
	CommonEntity // 使用 CommonEntity 对象代替 ID

	//ID            int64  `json:"id"`             // 视频唯一标识
	CreateDate 	int64 `json:"create_date,omitempty"`
	Author        User   `json:"author" gorm:"foreignKey:UserId;references:id"`                         // 视频作者信息
	UserId 	int64	`json:"user_id " gorm:"column:user_id" `
	CommentCount  int64  `json:"comment_count,omitempty"`            // 视频的评论总数
	CoverURL      string `json:"cover_url,omitempty"`                // 视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"`           // 视频的点赞总数
	IsFavorite    bool   `json:"is_favorite,omitempty"`              // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	Title         string `json:"title,omitempty"`                    // 视频标题
>>>>>>> Stashed changes
}
