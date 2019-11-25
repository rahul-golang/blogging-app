package models

// Followers *
type Followers struct {
	Times
	UserID     int `json:"user_id" `
	FollowerID int `json:"follower_id"`
}

//UserProfile *
type UserProfile struct {
	User      User        `json:"users"`
	Followers []Followers `json:""`
	Blogs     []Blog
}
