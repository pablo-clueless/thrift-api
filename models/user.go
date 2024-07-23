package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProps struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at"`
	Email     string             `json:"email" bson:"email"`
	ImageUrl  string             `json:"image_url" bson:"image_url"`
	LastLogin time.Time          `json:"last_login" bson:"last_login"`
	Name      string             `json:"name" bson:"name"`
	Password  string             `json:"password" bson:"password"`
	Streak    int                `json:"streak" bson:"streak"`
	Username  string             `json:"username" bson:"username"`
}
