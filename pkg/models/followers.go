package models

type Followers struct {
	Times
	UserID     int `json:"user_id" `
	FollowerID int `json:"follower_id"`
}

type UserProfile struct {
	User      User        `json:"users"`
	Followers []Followers `json:""`
	Blogs     []Blog
}
