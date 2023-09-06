package models

type Video struct {
	CommonEntity // 使用 CommonEntity 对象代替 ID
	//ID            int64  `json:"id"`             // 视频唯一标识
	AuthorID      int64  // 视频作者信息
	Author        User   `json:"author"`
	CommentCount  int64  `json:"comment_count,omitempty"`  // 视频的评论总数
	CoverURL      string `json:"cover_url,omitempty"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"` // 视频的点赞总数
	//IsFavorite    bool   `json:"is_favorite,omitempty"`              // true-已点赞，false-未点赞
	PlayURL string `json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	Title   string `json:"title,omitempty"`                    // 视频标题
}

func (table *Video) TableName() string {
	return "videos"
}
