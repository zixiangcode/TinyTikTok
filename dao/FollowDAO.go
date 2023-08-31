package dao

import (
	"TinyTikTok/models"
	"errors"
	"gorm.io/gorm"
)

// FollowUser 关注用户，实际上是数据库中添加一条关注关系
func FollowUser(db *gorm.DB, userId int64, toUserId int64) (err error) {
	follow := models.Follow{
		CommonEntity: models.CommonEntity{},
		UserId:       userId,
		FollowUserId: toUserId,
	}
	err = db.Create(&follow).Error
	return err
}

// UnFollowUser 取消关注，实际上将数据库中关注关系的 is_deleted 置为 1
func UnFollowUser(db *gorm.DB, userId int64, toUserId int64) (err error) {
	follow := models.Follow{
		UserId:       userId,
		FollowUserId: toUserId,
	}
	result := db.Model(&follow).Update("is_deleted", 1)
	if result.Error != nil {
		return result.Error
	}
	// 如果影响了 0 行数据，表明此关注关系不存在
	if result.RowsAffected == 0 {
		return errors.New("关注关系不存在")
	}

	return nil
}
