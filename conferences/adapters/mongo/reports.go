package mongo

import (
	"context"

	"github.com/VanillaFox/system_architecture_lab/conferences/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) CreateReport(ctx context.Context, report *models.Report) error {
	_, err := r.rep.InsertOne(ctx, report)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetReports(ctx context.Context, conferenceObjectID primitive.ObjectID) (models.Reports, error) {
	reports := make(models.Reports, 0)

	cur, err := r.rep.Find(ctx, bson.M{"conference_object_id": conferenceObjectID})

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var report models.ReportReadModel

		if err := cur.Decode(&report); err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func (r *Repository) UpdateReport(ctx context.Context, reportObjectID primitive.ObjectID, report models.Report) error {
	_, err := r.rep.UpdateOne(ctx, bson.M{"_id": reportObjectID}, bson.D{
		{"$set", bson.M{"title": report.Title}},
		{"$set", bson.M{"description": report.Description}}})

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteReport(ctx context.Context, reportObjectID primitive.ObjectID) error {
	_, err := r.rep.DeleteOne(ctx, bson.M{"_id": reportObjectID})

	if err != nil {
		return err
	}

	return nil
}
