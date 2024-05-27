package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConferenceCreateModel struct {
	Title string    `json:"title" bson:"title" example:"Конференция по JS"`
	Date  time.Time `json:"date" bson:"date" example:"2015-09-15T14:00:12-00:00"`
}

type ConferenceReadModel struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `json:"title" bson:"title" example:"Конференция по JS"`
	Date  time.Time          `json:"date" bson:"date" example:"2015-09-15T14:00:12-00:00"`
}

type Conferences []ConferenceReadModel
