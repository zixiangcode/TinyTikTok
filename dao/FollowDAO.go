package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

// GetFollowByUserIdAndToUserId 通过关注用户 id 和被关注用户 id 查询关注关系
func GetFollowByUserIdAndToUserId(userId int64, toUserId int64) (models.Follow, error) {
	result := models.Follow{}
	err := db.GetMysqlDB().Model(models.Follow{}).Where("user_id = ? AND follow_user_id = ?", userId, toUserId).Find(&result).Error
	return result, err
}

// FollowUser 关注用户，实际上是数据库中添加一条关注关系
func FollowUser(db *gorm.DB, userId int64, toUserId int64) (err error) {
	follow, err := GetFollowByUserIdAndToUserId(userId, toUserId)
	if err != nil {
		log.Printf("查询关注关系失败")
		return err
	}
	// 如果已经存在此关注关系，则将 is_deleted 置为 0
	if follow.Id != 0 {
		err = db.Model(&models.Follow{}).Where("user_id = ? AND follow_user_id = ?", userId, toUserId).Update("is_deleted", 0).Error
		if err != nil {
			log.Printf("更新关注关系失败")
			return err
		}
	} else {
		follow = models.Follow{
			CommonEntity: models.NewCommonEntity(),
			UserId:       userId,
			FollowUserId: toUserId,
		}
		if err := db.Create(&follow).Error; err != nil {
			log.Printf("添加关注关系失败: %s", err.Error())
			return err
		}
	}

	return err
}

// UnFollowUser 取消关注，实际上将数据库中关注关系的 is_deleted 置为 1
func UnFollowUser(db *gorm.DB, userId int64, toUserId int64) (err error) {
	follow := models.Follow{
		UserId:       userId,
		FollowUserId: toUserId,
	}
	result := db.Model(&follow).Where(&follow).UpdateColumn("is_deleted", 1)
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
	if len(followedUsers) > 0 {
		err = db.GetMysqlDB().Model(&models.User{}).Where("id NOT IN ?", followedUsers).Update("is_follow", 0).Error
		if err != nil {
			return err
		}
	}
	err = db.GetMysqlDB().Model(&models.User{}).Where("id IN ?", followedUsers).Update("is_follow", 1).Error
	if err != nil {
		return err
	}
	return nil
}

// TODO 点击取消关注，user 表的 is_follow 属性更新存在问题

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
	err = db.GetMysqlDB().Model(&models.User{}).Where("id IN ?", followedUsers).Update("is_follow", 1).Error
	if err != nil {
		return err
	}
	if len(followedUsers) > 0 {
		err = db.GetMysqlDB().Model(&models.User{}).Where("id NOT IN ?", followedUsers).Update("is_follow", 0).Error
		if err != nil {
			return err
		}
	}
	return nil
}
