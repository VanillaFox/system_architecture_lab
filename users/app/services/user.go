package services

import (
	"context"

	"github.com/VanillaFox/system_architecture_lab/users/adaptres/postgres"
	"github.com/VanillaFox/system_architecture_lab/users/models"
)

type UserService struct {
	repository *postgres.Repository
}

func NewUserService(repository *postgres.Repository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	return s.repository.CreateUser(ctx, user)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.repository.GetByUsername(ctx, username)
}

func (s *UserService) GetUserByFullNamePrefix(ctx context.Context, fullNamePrefix string) (*models.Users, error) {
	return s.repository.GetByFullNamePrefix(ctx, fullNamePrefix)
}

func (s *UserService) GetUsers(ctx context.Context) (models.Users, error) {
	return s.repository.GetUsers(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, username string, user *models.User) (*models.User, error) {
	return s.repository.UpdateUser(ctx, username, user)
}

func (s *UserService) DeleteUser(ctx context.Context, username string) error {
	return s.repository.DeleteUser(ctx, username)
}
