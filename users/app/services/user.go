package services

import (
	"context"

	"github.com/VanillaFox/system_architecture_lab/users/adaptres/cache"
	"github.com/VanillaFox/system_architecture_lab/users/adaptres/postgres"
	"github.com/VanillaFox/system_architecture_lab/users/models"
)

type UserService struct {
	repository *postgres.Repository
	cache      *cache.Cache
}

func NewUserService(repository *postgres.Repository, cache *cache.Cache) *UserService {
	return &UserService{
		repository: repository,
		cache:      cache,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	return s.cache.FirstSetUser(ctx, user)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.cache.GetUser(ctx, username)
}

func (s *UserService) GetUserByFullNamePrefix(ctx context.Context, fullNamePrefix string) (*models.Users, error) {
	return s.repository.GetByFullNamePrefix(ctx, fullNamePrefix)
}

func (s *UserService) GetUsers(ctx context.Context) (models.Users, error) {
	return s.repository.GetUsers(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, username string, user *models.User) (*models.User, error) {
	return s.cache.SetUser(ctx, username, user)
}

func (s *UserService) DeleteUser(ctx context.Context, username string) error {
	return s.repository.DeleteUser(ctx, username)
}
