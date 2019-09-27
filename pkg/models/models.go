package models

import (
	"time"
)

type Model struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"user_email"`
	Phone     string `json:"user_phone"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Blogs     []Blog `json:"blogs" gorm:"foreignkey:ID"`
}

type GetUserResp struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type CreateUserReq struct {
	User User `json:"user"`
}

type CreateUserResp struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type DeleteUserResp struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

type UpdateUserReq struct {
	User User `json:"user"`
}

type UpdateUserResp struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type GetAllUserResp struct {
	Message string  `json:"message"`
	User    []*User `json:"user"`
}

type Blog struct {
	Model
	Tittle    string `json:"tittle"`
	RelatedTo string `json:"related_to"`
	Containt  string `json:"containt"`
	UserID    int    `json:"user_id" gorm:"foreignkey:user_id"`
}

type CreateBlogReq struct {
	Blog Blog `json:"blog"`
}

type CreateBlogResp struct {
	Message string `json:"message"`
	Blog    *Blog  `json:"blog"`
}

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
