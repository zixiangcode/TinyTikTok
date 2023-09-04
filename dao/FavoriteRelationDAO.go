package dao

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
)

func AddfavoriteRelation(favoriteRelation models.FavoriteRelation) error {
	err := db.GetMysqlDB().Create(&favoriteRelation).Error
	return err
}

func DeletefavoriteRelation(favoriteRelation models.FavoriteRelation) error {

	//直接删除数据库数据，因为用户可以频繁点赞再取消点赞，会产生大量数据，如果在点赞时先判断数据库中是否有记录，再做修改或者新增数据，操作更加繁琐
	err := db.GetMysqlDB().Model(models.FavoriteRelation{}).
		Where("user_id = ? && video_id = ?", favoriteRelation.UserID, favoriteRelation.VideoID).
		Delete(favoriteRelation).
		Error
	return err
}
