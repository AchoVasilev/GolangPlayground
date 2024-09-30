package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	Id        primitive.ObjectID `bson:"_id"`
	MenuId    string             `json:"menu_id" validate:"required"`
	Name      string             `json:"name" validate:"required"`
	Category  string             `json:"category" validate:"required"`
	StartDate *time.Time         `json:"start_date" validate:"required"`
	EndDate   *time.Time         `json:"end_date" validate:"required"`
	CreatedAt time.Time          `json:"created_at" validate:"required"`
	UpdatedAt time.Time          `json:"updated_at"`
}
