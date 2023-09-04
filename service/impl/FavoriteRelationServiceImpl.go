package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
)

type FavoriteRelationServiceImpl struct {
}

func (favoriteRelationServiceImpl FavoriteRelationServiceImpl) AddfavoriteRelation(favoriteRelation models.FavoriteRelation) error {
	err := dao.AddfavoriteRelation(favoriteRelation)
	return err
}

func (favoriteRelationServiceImpl FavoriteRelationServiceImpl) DeletefavoriteRelation(favoriteRelation models.FavoriteRelation) error {
	err := dao.DeletefavoriteRelation(favoriteRelation)
	return err
}
