package service

import (
	"TinyTikTok/models"
)

type CommentService interface {
	Add(comment models.Comment) (int64, error)
	DeleteComment(commentID int64) error
}
