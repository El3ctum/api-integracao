package service

import "github.com/couchbase/gocb/v2"

type ServiceInitializaer struct {
	UserService *UserService
}

func InitServices(scope *gocb.Scope) *ServiceInitializaer {
	return &ServiceInitializaer{
		UserService: NewUserService(scope),
	}
}
