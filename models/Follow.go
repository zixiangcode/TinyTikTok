package models

// Follow 关注关系实体
type Follow struct {
	CommonEntity
	UserId       int64 `json:"user_id"`
	FollowUserId int64 `json:"follow_user_id"`
}

func (table *Follow) TableName() string {
	return "follow"
}

// FollowResponse 关注操作请求参数
type FollowResponse struct {
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

type FollowListResponse struct {
	StatusCode         int32  `json:"status_code"`
	StatusMsg          string `json:"status_msg"`
	UserFollowResponse []User `json:"user_list"`
}
