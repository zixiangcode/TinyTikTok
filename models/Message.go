package models

type Message struct {
	CommonEntity
	//ID         int64  `json:"id"`           // 消息id
	//CreateTime int64  `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID int64  `json:"from_user_id,omitempty"` // 消息发送者id
	Content    string `json:"content,omitempty"`      // 消息内容
	ToUserID   int64  `json:"to_user_id,omitempty"`   // 消息接收者id
}

func (table *Message) TableName() string {
	return "messages"
}
