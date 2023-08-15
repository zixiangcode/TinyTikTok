package models

type Comment struct {
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

type MyComment struct {
	Id         uint64 `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
