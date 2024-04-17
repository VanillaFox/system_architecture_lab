package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportCreateModel struct {
	Title       string `json:"title" example:"Доклад 1"`
	Description string `json:"description" example:"Первый докад"`
}

type Report struct {
	ConferenceObjectID primitive.ObjectID `bson:"conference_object_id"`
	Title              string             `bson:"title"`
	Description        string             `bson:"description"`
}

type ReportReadModel struct {
	ID                 primitive.ObjectID `bson:"_id" json:"id"`
	ConferenceObjectID primitive.ObjectID `bson:"conference_object_id" json:"conference_object_id"`
	Title              string             `json:"title" bson:"title" example:"Доклад 1"`
	Description        string             `json:"description" bson:"description" example:"Первый докад"`
}

type Reports []ReportReadModel
