package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Blog related information is here
type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Tittle    string             `json:"tittle" bson:"tittle"`
	RelatedTo string             `json:"related_to" bson:"related_to"`
	Containt  string             `json:"containt" bson:"containt"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Likes     int                `json:"likes" bson:"likes"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
