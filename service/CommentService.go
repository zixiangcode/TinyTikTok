package service

import (
	"TinyTikTok/models"
)

type CommentService interface {
	AddComment(comment models.Comment) (models.CommentCommonResponse, error)
	DeleteComment(commentID int64, userID int64) error
}
