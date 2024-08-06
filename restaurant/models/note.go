package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Id        primitive.ObjectID `bson:"_id"`
	Text      string             `json:"text" validate:"required,min=2,max=600"`
	Title     string             `json:"title" validate:"required,min=2,max=100"`
	CreatedAt time.Time          `json:"created_at" validate:"required"`
	UpdatedAt time.Time          `json:"updated_at"`
	NoteId    string             `json:"note_id" validate:"required"`
}
