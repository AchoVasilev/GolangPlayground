package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	FirstName    string             `json:"first_name" validate:"required,min=2,max=100"`
	LastName     string             `json:"last_name" validate:"required,min=2,max=100"`
	Email        string             `json:"email" validate:"required,email"`
	Password     string             `json:"password" validate:"required"`
	Avatar       string             `json:"avatar" validate:"required"`
	Phone        string             `json:"phone" validate:"required"`
	Token        string             `json:"token" validate:"required"`
	RefreshToken string             `json:"refresh_token" validate:"required"`
	CreatedAt    time.Time          `json:"created_at" validate:"required"`
	UpdatedAt    time.Time          `json:"updated_at" validate:"required"`
	UserId       string             `json:"user_id" validate:"required"`
}
