package models

//UserProfile *
type UserProfile struct {
	User      User        `json:"users"`
	Followers []Followers `json:""`
	Blogs     []Blog
}
