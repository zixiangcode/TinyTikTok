package models

type CommentListResponse struct {
	Response
	CommentList []CommentCommonResponse `json:"comment_list,omitempty"`
}
