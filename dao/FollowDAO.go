package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"errors"
	"gorm.io/gorm"
)

// FollowUser 关注用户，实际上是数据库中添加一条关注关系
func FollowUser(db *gorm.DB, userId int64, toUserId int64) (err error) {
	follow := models.Follow{
		CommonEntity: models.NewCommonEntity(),
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

// todo 热更新前端显示

// todo 取消关注更新数据库

// UpdateUserFollowByUserName 通过用户名更新 user 表的 is_follow 属性列
func UpdateUserFollowByUserName(username string) (err error) {
	// 查询对应用户名对应的 userId
	var userId int64
	err = db.GetMysqlDB().Model(&models.User{}).Select("id").Where("name = ?", username).First(&userId).Error
	if err != nil {
		return err
	}
	// 查询关注关系
	var followedUsers []int64
	err = db.GetMysqlDB().Table("follow").
		Where("user_id = ? AND is_deleted = ?", userId, 0).
		Pluck("follow_user_id", &followedUsers).Error
	if err != nil {
		return err
	}
	// 更新 is_follow 属性列
	err = db.GetMysqlDB().Model(&models.User{}).Where("id NOT IN ?", followedUsers).Update("is_follow", 0).Error
	if err != nil {
		return err
	}
	err = db.GetMysqlDB().Model(&models.User{}).Where("id IN ?", followedUsers).Update("is_follow", 1).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserFollowByUserId 通过用户 ID 更新 user 表的 is_follow 属性列
func UpdateUserFollowByUserId(userId int64) (err error) {
	// 查询关注关系
	var followedUsers []int64
	err = db.GetMysqlDB().Table("follow").
		Where("user_id = ? AND is_deleted = ?", userId, 0).
		Pluck("follow_user_id", &followedUsers).Error
	if err != nil {
		return err
	}
	// 更新 is_follow 属性列
	err = db.GetMysqlDB().Model(&models.User{}).Where("id NOT IN ?", followedUsers).Update("is_follow", 0).Error
	if err != nil {
		return err
	}
	err = db.GetMysqlDB().Model(&models.User{}).Where("id IN ?", followedUsers).Update("is_follow", 1).Error
	if err != nil {
		return err
	}
	return nil
}
