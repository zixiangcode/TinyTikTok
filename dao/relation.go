package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"errors"
)

func AddFollow(owner *models.User, Target *models.User) error {
	relation := models.Relation{}
	tx := db.GORM.Begin()
	OwnerId := owner.Id
	TargetId := Target.Id

	relation.OwnerId = OwnerId
	relation.TargetID = TargetId
	if t := db.GORM.Where("owner_id = ? and target_id = ?", OwnerId, TargetId).Find(&relation); t.RowsAffected != 0 {
		tx.Rollback()
		return errors.New("该关系已存在")
	}
	if t := tx.Create(&relation); t.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("关注失败")
	}

	tx.Commit()
	return nil
}

func GetFollow(id int64) *[]models.User {
	relation := make([]models.Relation, 0)
	db.GORM.Where("owner_id = ? ", id).Find(&relation)
	userID := make([]int64, 0)

	for _, v := range relation {
		userID = append(userID, v.TargetID)
	}

	user := make([]models.User, 0)
	db.GORM.Where("id in ?", userID).Find(&user)
	user2 := make([]models.User, len(user))
	for i, v := range user {
		user2[i] = v
		user2[i].IsFollow = true
	}

	return &user2
}

func GetFollower(id int64) *[]models.User {
	relation := make([]models.Relation, 0)
	db.GORM.Where("target_id = ? ", id).Find(&relation)
	userID := make([]int64, 0)

	for _, v := range relation {
		userID = append(userID, v.OwnerId)
	}

	user := make([]models.User, 0)
	db.GORM.Where("id in ?", userID).Find(&user)

	return &user
}

func GetFriend(id int64) *[]models.User {
	relation := make([]models.Relation, 0)
	db.GORM.Where("target_id = ? or owner_id = ?", id, id).Find(&relation)
	userID := make([]int64, 0)

	for _, v := range relation {
		userID = append(userID, v.OwnerId)
	}

	user := make([]models.User, 0)
	db.GORM.Where("id in ?", userID).Find(&user)

	return &user
}
