package mongo

import (
	"context"

	"github.com/VanillaFox/system_architecture_lab/conferences/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) CreateConference(ctx context.Context, conference models.ConferenceCreateModel) error {
	_, err := r.conf.InsertOne(ctx, conference)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetConferences(ctx context.Context) (models.Conferences, error) {
	conferences := make(models.Conferences, 0)

	cur, err := r.conf.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var conference models.ConferenceReadModel

		if err := cur.Decode(&conference); err != nil {
			return nil, err
		}

		conferences = append(conferences, conference)
	}

	return conferences, nil
}

func (r *Repository) UpdateConference(ctx context.Context, objectID primitive.ObjectID, conference models.ConferenceCreateModel) error {
	_, err := r.conf.UpdateOne(ctx, bson.M{"_id": objectID}, bson.D{{"$set", conference}})

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteConference(ctx context.Context, objectID primitive.ObjectID) error {
	_, err := r.conf.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	_, err = r.rep.DeleteMany(ctx, bson.M{"conference_object_id": objectID})
	if err != nil {
		return err
	}

	return nil
}
