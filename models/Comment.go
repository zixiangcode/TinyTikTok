package models

type Comment struct {
	CommonEntity // 此中已包含了 CreateTime
	//ID         int64  `json:"id"`          // 评论id
	Content string `json:"content"` // 评论内容
	//CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	User User `json:"user"` // 评论用户信息
}
