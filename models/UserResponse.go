package models

type UserResponse struct {
	Response
	User User `json:"user"`
}
