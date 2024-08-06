package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	Id             primitive.ObjectID `bson:"_id"`
	NumberOfGuests int                `json:"number_of_guests"`
	CreatedAt      time.Time          `json:"created_at" validate:"required"`
	UpdatedAt      time.Time          `json:"updated_at"`
	TableId        string             `json:"table_id" validate:"required"`
}
