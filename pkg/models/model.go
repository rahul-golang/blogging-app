package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
