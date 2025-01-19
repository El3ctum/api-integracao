package service

import (
	"api-integracao/internal/models"

	"github.com/couchbase/gocb/v2"
)

type IUserService interface {
	CreateUser(string, *models.User) error
	UpdateUser(string, *models.User) error
	DeleteUser(string) error
	GetUserById(string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	ListUsers(string, int, int) ([]models.User, error)
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
	_, err := s.scope.Collection(s.collectionName).Insert(docKey, data, nil)
	return err
}

func (s *UserService) UpdateUser(docKey string, data *models.User) error {
	_, err := s.scope.Collection(s.collectionName).Upsert(docKey, data, nil)
	return err
}

func (s *UserService) DeleteUser(docKey string) error {
	_, err := s.scope.Collection(s.collectionName).Remove(docKey, nil)
	return err
}

func (s *UserService) GetUserById(docKey string) (*models.User, error) {
	var user models.User
	query := `SELECT u.* FROM ` + s.collectionName + ` u WHERE META(u).id = $id`
	queryResult, err := s.scope.Query(query, &gocb.QueryOptions{
		NamedParameters: map[string]interface{}{
			"id": docKey,
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

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User

	n1qlQuery := `SELECT u.* FROM ` + s.collectionName + ` u`
	rows, err := s.scope.Query(n1qlQuery, nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Row(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) ListUsers(query string, limit int, offset int) ([]models.User, error) {
	var users []models.User

	n1qlQuery := `SELECT u.* FROM ` + s.collectionName + ` u WHERE u.name LIKE $query LIMIT $limit OFFSET $offset`
	queryResult, err := s.scope.Query(n1qlQuery, &gocb.QueryOptions{
		NamedParameters: map[string]interface{}{
			"query":  "%" + query + "%",
			"limit":  limit,
			"offset": offset,
		},
	})
	if err != nil {
		return nil, err
	}
	defer queryResult.Close()

	for queryResult.Next() {
		var user models.User
		err := queryResult.Row(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := queryResult.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
