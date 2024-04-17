package services

import (
	"context"

	mongoRepo "github.com/VanillaFox/system_architecture_lab/conferences/adapters/mongo"
	"github.com/VanillaFox/system_architecture_lab/conferences/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportService struct {
	repository *mongoRepo.Repository
}

func NewReportService(repository *mongoRepo.Repository) *ReportService {
	return &ReportService{
		repository: repository,
	}
}

func (s *ReportService) CreateReport(ctx context.Context, report *models.Report) error {
	return s.repository.CreateReport(ctx, report)
}

func (s *ReportService) GetReports(ctx context.Context, conferenceObjectID primitive.ObjectID) (models.Reports, error) {
	return s.repository.GetReports(ctx, conferenceObjectID)
}

func (s *ReportService) UpdateReport(ctx context.Context, reportObjectID primitive.ObjectID, report *models.Report) error {
	return s.repository.UpdateReport(ctx, reportObjectID, *report)
}

func (s *ReportService) DeleteReport(ctx context.Context, reportObjectID primitive.ObjectID) error {
	return s.repository.DeleteReport(ctx, reportObjectID)
}
