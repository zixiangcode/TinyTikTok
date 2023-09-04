package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"gorm.io/gorm"
)

func AddComment(comment models.Comment) (int64, error) {
	// 开始事务
	tx := db.GetMysqlDB().Begin()

	// 创建评论记录
	result := tx.Create(&comment)
	if result.Error != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return 0, result.Error
	}

	// 更新视频的评论总数
	if err := tx.Model(models.Video{}).
		Where("id = ?", comment.VideoID).
		Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return 0, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return comment.Id, nil
}

func DeleteComment(commentID int64, userID int64, videoID int64) error {
	// 开始事务
	tx := db.GetMysqlDB().Begin()

	// 标记评论为已删除
	if err := tx.Model(&models.Comment{}).
		Where("id = ? AND user_id = ?", commentID, userID).
		Update("is_deleted", 1).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新视频的评论总数，减去一个评论
	if err := tx.Model(models.Video{}).
		Where("id = ?", videoID).
		Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func GetCommentsByVideoID(videoID int64) ([]models.Comment, error) {
	var comments []models.Comment
	result := db.GetMysqlDB().
		Where("video_id = ? AND is_deleted <> 1", videoID).
		Order("create_time desc").
		Find(&comments)
	return comments, result.Error
}

func GetCommentByID(commentID int64) (models.Comment, error) {
	var comment models.Comment
	err := db.GetMysqlDB().Model(models.Comment{}).
		Where("id = ? AND is_deleted <> 1", commentID).
		Find(&comment).Error
	return comment, err
}
