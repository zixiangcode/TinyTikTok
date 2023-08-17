package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
)

func AddComment(comment models.Comment) (int64, error) {

	result := db.GetMysqlDB().Create(&comment)
	return comment.Id, result.Error
}

func DeleteComment(commentID int64) error {
	result := db.GetMysqlDB().Model(&models.Comment{}).
		Where("id = ?", commentID).
		Update("is_deleted", 1)
	return result.Error
}

func GetCommentsByVideoID(videoID int64) ([]models.Comment, error) {
	var comments []models.Comment
	result := db.GetMysqlDB().Preload("User").
		Find(&comments, "video_id = ?", videoID).
		Order("create_date desc")
	return comments, result.Error
}
