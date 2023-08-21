package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"errors"
)

func FindUserById(id int64) (*models.User, error) {
	user := models.User{}
	if tx := db.GORM.Where("id = ?", id).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("no user")
	}
	return &user, nil
}

func FindUserByToken(token string) (*models.User, error) {
	user := models.User{}
	if tx := db.GORM.Where("token = ?", token).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("no user")
	}
	return &user, nil
}
