package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionProps struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt   time.Time          `json:"deleted_at" bson:"deleted_at"`
	AccountId   string             `json:"account_id" bson:"account_id"`
	Amount      float64            `json:"amount" bson:"amount"`
	Category    string             `json:"category" bson:"category"`
	Date        time.Time          `json:"date" bson:"date"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	Type        string             `json:"type" bson:"type"`
}
