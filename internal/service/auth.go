package service

import (
	"api-integracao/internal/models"

	"github.com/couchbase/gocb/v2"
)

type IAuthService interface {
	GetUserByEmail(*models.User) (*models.User, error)
	RegisterUser(*models.User) error
}

type AuthService struct {
	collectionName string
	scope          *gocb.Scope
}

func NewAuthService(scope *gocb.Scope) *AuthService {
	return &AuthService{
		collectionName: "users",
		scope:          scope,
	}
}

func (s *AuthService) GetUserByEmail(user *models.User) (*models.User, error) {
	return nil, nil
}

func (s *AuthService) RegisterUser(*models.User) error {
	return nil
}
