package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User **
type User struct {
	Times
	ID        primitive.ObjectID `json:"id" bson:"id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Email     string             `json:"user_email" bson:"user_email"`
	Phone     string             `json:"user_phone" bson:"user_phone"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
}

// Followers *
type Followers struct {
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id" `
	FollowerID primitive.ObjectID `json:"follower_id" bson:"follower_id"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt  *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
