package service

import "TinyTikTok/models"

type UserFollowService interface {
	Add(comment models.UserFollow) error
}
