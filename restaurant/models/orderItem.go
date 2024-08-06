package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	Id          primitive.ObjectID `bson:"_id"`
	Quantity    string             `json:"quantity" validate:"required,eq=SMALL|MEDIUM|LARGE"`
	UnitPrice   float64            `json:"unit_price" validate:"required"`
	CreatedAt   time.Time          `json:"created_at" validate:"required"`
	UpdatedAt   time.Time          `json:"updated_at" validate:"required"`
	FoodId      string             `json:"food_id" validate:"required"`
	OrderId     string             `json:"order_id" validate:"required"`
	OrderItemId string             `json:"order_item_id" validate:"required"`
}
