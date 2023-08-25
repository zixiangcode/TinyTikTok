package service

import (
	"TinyTikTok/models"
)

type FollowService interface {
	AddFollowAction(userID int64, to_userID int64) error

	DelFollowAction(userID int64, to_userID int64) error

	GetFollowList(Id int64) *[]models.User

	GetFollowerList(Id int64) *[]models.User

	GetFriendList(Id int64) *[]models.User
}
