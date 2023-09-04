package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"gorm.io/gorm"
)

func AddfavoriteRelation(favoriteRelation models.FavoriteRelation) error {
	// 开始事务
	tx := db.GetMysqlDB().Begin()

	// 创建点赞关系记录
	if err := tx.Create(&favoriteRelation).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新用户的 favorite_count
	if err := tx.Model(models.User{}).
		Where("id = ?", favoriteRelation.UserID).
		Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 查找视频信息
	var video models.Video
	if err := tx.
		Where("id = ?", favoriteRelation.VideoID).
		Find(&video).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新用户的 total_favorited
	if err := tx.Model(models.User{}).
		Where("id = ?", video.AuthorID).
		Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新video中的获赞总数
	if err := tx.Model(models.Video{}).
		Where("id = ?", favoriteRelation.VideoID).
		Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

func DeletefavoriteRelation(favoriteRelation models.FavoriteRelation) error {

	// 开始事务
	tx := db.GetMysqlDB().Begin()

	// 删除点赞关系记录
	if err := tx.Where("user_id = ? && video_id = ?", favoriteRelation.UserID, favoriteRelation.VideoID).
		Delete(models.FavoriteRelation{}).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新用户的 favorite_count
	if err := tx.Model(models.User{}).
		Where("id = ?", favoriteRelation.UserID).
		Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 查找视频信息
	var video models.Video
	if err := tx.
		Where("id = ?", favoriteRelation.VideoID).
		Find(&video).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新用户的 total_favorited
	if err := tx.Model(models.User{}).
		Where("id = ?", video.AuthorID).
		Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新video中的获赞总数
	if err := tx.Model(models.Video{}).
		Where("id = ?", favoriteRelation.VideoID).
		Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

func IsfavoriteRelation(userID int64, videoID int64) (models.FavoriteRelation, error) {
	var favoriteRelation models.FavoriteRelation
	err := db.GetMysqlDB().
		Where("user_id = ? AND video_id = ?", userID, videoID).Find(&favoriteRelation).Error
	return favoriteRelation, err
}
