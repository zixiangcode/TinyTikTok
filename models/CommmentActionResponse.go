package models

type CommentActionResponse struct {
	Response
	Comment MyComment `json:"comment,omitempty"`
}
