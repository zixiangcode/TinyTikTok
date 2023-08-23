package models

type CommentActionResponse struct {
	StatusCode int32                 `json:"status_code"`
	StatusMsg  string                `json:"status_msg,omitempty"`
	Comment    CommentCommonResponse `json:"comment,omitempty"`
}
