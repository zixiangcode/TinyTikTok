package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
)

func SendMessage(message models.Message) error {
	result := db.GetMysqlDB().Create(&message)
	return result.Error
}

func GetMessages(userID int64, toUserId int64) ([]models.Message, error) {
	var messages []models.Message
	result := db.GetMysqlDB().
		Where("(from_user_id = ? AND to_user_id = ? ) OR (from_user_id = ? AND to_user_id = ? )", userID, toUserId, toUserId, userID).
		Order("create_time asc").
		Find(&messages)
	return messages, result.Error
}
