package models

type User struct {
	Model
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"user_email" bson:"user_email"`
	Phone     string `json:"user_phone" bson:"user_phone"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Blogs     []Blog `json:"blogs" bson:"blogs" gorm:"foreignkey:ID"`
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
