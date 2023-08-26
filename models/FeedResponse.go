package models

type FeedResponse struct {
	StatusCode int64                   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string                  `json:"status_msg"`  // 返回状态描述
	NextTime   int64                   `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList  []FeedResponseVideoInfo `json:"video_list"`  // 视频列表
}

// Video
type FeedResponseVideoInfo struct {
	ID            int64        `json:"id"`             // 视频唯一标识
	Author        FeedUserInfo `json:"author"`         // 视频作者信息
	PlayURL       string       `json:"play_url"`       // 视频播放地址
	CoverURL      string       `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64        `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64        `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool         `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string       `json:"title"`          // 视频标题
}

// 视频作者信息
//
// User
type FeedUserInfo struct {
	ID              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
}

// 从数据中查询的视频数据格式
type FeedVideoInfo struct {
	AuthorID      int64
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}
