package service

import "TinyTikTok/models"

type FavoriteRelationService interface {
	AddfavoriteRelation(favoriteRelation models.FavoriteRelation) error
	DeletefavoriteRelation(favoriteRelation models.FavoriteRelation) error
}
