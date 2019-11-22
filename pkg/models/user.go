package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
