package models

type CommentActionResponse struct {
	Response
	Comment CommentCommonResponse `json:"comment,omitempty"`
}
