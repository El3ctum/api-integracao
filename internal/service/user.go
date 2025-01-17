package service

import (
	"api-integracao/internal/models"
	"fmt"

	"github.com/couchbase/gocb/v2"
)

type IUserService interface {
	CreateUser(string, *models.User) error
	// GetUser(string) (*models.User, error)
	// UpdateUser(string, *models.User) error
	// DeleteUser(string) error
	// ListUsers(string, int, int) ([]models.User, error)
}

type UserService struct {
	collectionName string
	scope          *gocb.Scope
}

func NewUserService(scope *gocb.Scope) *UserService {
	return &UserService{
		collectionName: "users",
		scope:          scope,
	}
}

func (s *UserService) CreateUser(docKey string, data *models.User) error {
	fmt.Println("Criando Usuário...")
	_, err := s.scope.Collection(s.collectionName).Insert(docKey, data, nil)
	if err != nil {
		return err
	}
	return nil
}
