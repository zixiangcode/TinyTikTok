package service

import "TinyTikTok/models"

type MessageService interface {
	// SendMessage 发送消息
	SendMessage(toUserid string, fromUserId int64, content string) error
	// GetMessages 读取消息列
	GetMessages(userID int64, toUserid string) ([]models.MessageResponse, error)
}
