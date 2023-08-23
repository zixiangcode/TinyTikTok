package models

import "time"

//UserFollow 关注关系实体
type UserFollow struct {
	CommonEntity
	UserID       int64     `json:"user_id"`        // 用户id
	FollowUserID int64     `json:"follow_user_id"` // 关注用户ID
	UpdateTime   time.Time `json:"create_date"`    // 更新时间
	ActionType   int       `json:"action_type"`    // 1-关注，2-取消关注
}

func (table *UserFollow) TableName() string {
	return "user_follow"
}

//UserFollowResp 关注操作请求参数
type UserFollowResp struct {
	ID              int64  `json:"id"`               // 关注用户ID
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        //
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorite   int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
}

type UserFollowListResponse struct {
	StatusCode     int32            `json:"status_code"`
	StatusMsg      string           `json:"status_msg"`
	UserFollowResp []UserFollowResp `json:"user_list"`
}
