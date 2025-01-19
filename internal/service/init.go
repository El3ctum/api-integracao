package service

import "github.com/couchbase/gocb/v2"

type ServiceInitializaer struct {
	AuthService *AuthService
	UserService *UserService
}

func InitServices(scope *gocb.Scope) *ServiceInitializaer {
	return &ServiceInitializaer{
		AuthService: NewAuthService(scope),
		UserService: NewUserService(scope),
	}
}
