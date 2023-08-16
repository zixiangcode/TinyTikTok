package service

import (
	"TinyTikTok/models"
)

type CommentService interface {
	Add(comment models.Comment) (int64, error)
	Delete(commentID int64) error
}
