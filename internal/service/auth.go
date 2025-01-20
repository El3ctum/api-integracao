package service

import (
	"api-integracao/internal/models"

	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
)

type IAuthService interface {
	GetUserByEmail(string) (*models.User, error)
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

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT u.* FROM ` + s.collectionName + ` u WHERE u.email = $email`
	queryResult, err := s.scope.Query(query, &gocb.QueryOptions{
		NamedParameters: map[string]interface{}{
			"email": email,
		},
	})
	if err != nil {
		return nil, err
	}
	defer queryResult.Close()

	if queryResult.Next() {
		err := queryResult.Row(&user)
		if err != nil {
			return nil, err
		}
	}

	if err := queryResult.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) RegisterUser(data *models.User) error {
	uuid := uuid.New()
	docKey := uuid.String()
	_, err := s.scope.Collection(s.collectionName).Insert(docKey, data, nil)
	return err
}
