package models

type Comment struct { //储存在数据库中的评论结构
	CommonEntity // 此中已包含了 CreateTime
	//ID         int64  `json:"id"`          // 评论id
	Content string `json:"content"` // 评论内容
	//CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	//User User `json:"user"` // 评论用户信息
	UserID  int64 `json:"user_id"`
	VideoID int64 `json:"video_id"`
}

func (table *Comment) TableName() string {
	return "comments"
}

type CommentUserInfo struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type CommentCommonResponse struct { //返回的评论信息结构体
	Id         int64           `json:"id,omitempty"`
	User       CommentUserInfo `json:"user"`
	Content    string          `json:"content,omitempty"`
	CreateTime string          `json:"create_date,omitempty"`
}
