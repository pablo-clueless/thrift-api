package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountProps struct {
	Id           primitive.ObjectID  `json:"id" bson:"_id"`
	CreatedAt    time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at" bson:"updated_at"`
	DeletedAt    time.Time           `json:"deleted_at" bson:"deleted_at"`
	Balance      float64             `json:"balance" bson:"balance"`
	Name         string              `json:"name" bson:"name"`
	Transactions []*TransactionProps `json:"transactions,omitempty" bson:"transactions,omitempty"`
	UserId       string              `json:"user_id" bson:"user_id"`
}
