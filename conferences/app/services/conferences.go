package services

import (
	"context"

	mongoRepo "github.com/VanillaFox/system_architecture_lab/conferences/adapters/mongo"
	"github.com/VanillaFox/system_architecture_lab/conferences/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConferenceService struct {
	repository *mongoRepo.Repository
}

func NewConferenceService(repository *mongoRepo.Repository) *ConferenceService {
	return &ConferenceService{
		repository: repository,
	}
}

func (s *ConferenceService) CreateConference(ctx context.Context, conference *models.ConferenceCreateModel) error {
	return s.repository.CreateConference(ctx, *conference)
}

func (s *ConferenceService) GetConferences(ctx context.Context) (models.Conferences, error) {
	return s.repository.GetConferences(ctx)
}

func (s *ConferenceService) DeleteConference(ctx context.Context, objectID primitive.ObjectID) error {
	return s.repository.DeleteConference(ctx, objectID)
}

func (s *ConferenceService) UpdateConference(ctx context.Context, objectID primitive.ObjectID, confernce *models.ConferenceCreateModel) error {
	return s.repository.UpdateConference(ctx, objectID, *confernce)
}
