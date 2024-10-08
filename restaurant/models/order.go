package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id        primitive.ObjectID `bson:"_id"`
	OrderDate time.Time          `json:"order_date" validate:"required"`
	CreatedAt time.Time          `json:"created_at" validate:"required"`
	UpdatedAt time.Time          `json:"updated_at"`
	OrderId   string             `json:"order_id" validate:"required"`
}
