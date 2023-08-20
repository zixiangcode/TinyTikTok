package models
type User struct {
	Id           		int64   `json:"id,omitempty"`
	Name         		string  `json:"name,omitempty"`
	FollowCount   		int64   `json:"follow_count,omitempty"`
	FollowerCount 		int64   `json:"follower_count,omitempty"`
	IsFollow      		bool    `json:"is_follow,omitempty"`
	Avatar 		 		string	`json:"avatar,omitempty"`
	Background_image 	string	`json:"background_image,omitempty"`
	Signature			string	`json:"signature,omitempty"`
	Total_favorited		string	`json:"total_favorited,omitempty"`
	Work_count			int64	`json:"work_count,omitempty"`
	Favorite_count		int64	`json:"favorite_count,omitempty"`
}
