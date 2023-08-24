package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
)

func AddComment(comment models.Comment) (int64, error) {

	result := db.GetMysqlDB().Create(&comment)
	return comment.Id, result.Error
}

func DeleteComment(commentID int64, userID int64) error {
	result := db.GetMysqlDB().Model(&models.Comment{}).
		Where("id = ? AND user_id = ?", commentID, userID).
		Update("is_deleted", 1)
	return result.Error
}

func GetCommentsByVideoID(videoID int64) ([]models.Comment, error) {
	var comments []models.Comment
	result := db.GetMysqlDB().
		Where("video_id = ? AND is_deleted <> 1", videoID).
		Order("create_time desc").
		Find(&comments)
	return comments, result.Error
}
