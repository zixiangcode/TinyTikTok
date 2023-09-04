package service

import "TinyTikTok/models"

type RelationService interface {
	// FollowUser 关注用户
	FollowUser(userId int64, toUserId int64, actionType int) error
	// GetFollows 获取自身关注列表
	GetFollows(userId int64) ([]models.User, error)
	// GetFollowers 获取粉丝列表
	GetFollowers(userId int64) ([]models.User, error)
	// GetFriends 获取好友列表
	GetFriends(userId int64) ([]models.User, error)
}
