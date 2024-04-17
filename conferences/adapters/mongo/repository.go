package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	conf *mongo.Collection
	rep  *mongo.Collection
}

func NewRepository(conf *mongo.Collection, rep *mongo.Collection) *Repository {
	return &Repository{conf: conf, rep: rep}
}
