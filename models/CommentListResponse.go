package models

type CommentListResponse struct {
	StatusCode  int32                   `json:"status_code"`
	StatusMsg   string                  `json:"status_msg,omitempty"`
	CommentList []CommentCommonResponse `json:"comment_list,omitempty"`
}
