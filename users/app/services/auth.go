package services

import (
	"context"
	"errors"

	"github.com/VanillaFox/system_architecture_lab/users/adaptres/postgres"
	"github.com/VanillaFox/system_architecture_lab/users/models"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
)

type AuthService struct {
	repository   *postgres.Repository
	jwtSecretKey []byte
}

var ErrInvalidUsernameOrPassword = errors.New("invalid username or password")

func NewAuthService(repository *postgres.Repository, jwtSecretKey []byte) *AuthService {
	return &AuthService{
		repository:   repository,
		jwtSecretKey: jwtSecretKey,
	}
}

func (s *AuthService) Auth(ctx context.Context, creds *models.Creds) (models.AuthResponse, error) {
	password, err := s.repository.Auth(ctx, creds)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.AuthResponse{}, ErrInvalidUsernameOrPassword
		}

		return models.AuthResponse{}, err
	}

	passwordIsValid := creds.Password.Check(password)

	if !passwordIsValid {
		return models.AuthResponse{}, ErrInvalidUsernameOrPassword
	}

	payload := jwt.MapClaims{
		"username": creds.Username,
		"password": creds.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(s.jwtSecretKey)

	if err != nil {
		return models.AuthResponse{}, err
	}

	return models.AuthResponse{Token: signedToken}, err
}
