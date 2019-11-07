package models

type Followers struct {
	Model
	UserID     int `json:"user_id" `
	FollowerID int `json:"follower_id"`
}

type UserProfile struct {
	User      User        `json:"users"`
	Followers []Followers `json:""`
	Blogs     []Blog
}
