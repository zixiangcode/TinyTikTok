package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"gorm.io/gorm"
)

func AddUserFollow(db *gorm.DB, userFollow models.UserFollow) (err error) {
	err = db.Create(&userFollow).Error
	return
}

func GetUserFollowBy(db *gorm.DB, userID, followUserID int64) (userFollow models.UserFollow, err error) {
	err = db.Where("user_id = ? and follow_user_id  =? and is_deleted = ? ", userID, followUserID, 0).First(&userFollow).Error
	return
}

func GetUserFollows(userID int64) (userFollows []models.UserFollow, err error) {
	err = db.GetMysqlDB().Where("user_id = ? and action_type = ? and is_deleted = ?", userID, 1, 0).Find(&userFollows).Error
	return
}

func UpdateFollow(db *gorm.DB, id int64, userFollow models.UserFollow) (err error) {
	err = db.Where("id = ? and is_deleted = ?", id, 0).Updates(&userFollow).Error
	return
}
