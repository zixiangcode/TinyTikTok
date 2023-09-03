package models

type MessageResponse struct {
	ID         int64  `json:"id"`           // 消息id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	Content    string `json:"content"`      // 消息内容
	CreateTime string `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
}

type MessageResponseList struct {
	StatusCode          int32             `json:"status_code"`
	StatusMsg           string            `json:"status_msg,omitempty"`
	MessageResponseList []MessageResponse `json:"message_list"` // 用户列表
}
