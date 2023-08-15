package models

type CommentListResponse struct {
	Response
	CommentList []MyComment `json:"comment_list,omitempty"`
}
