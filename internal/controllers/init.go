package controllers

import service "api-integracao/internal/service"

type ControllerInitializaer struct {
	UserController *UserController
}

func InitControllers(services service.ServiceInitializaer) ControllerInitializaer {
	return ControllerInitializaer{
		UserController: NewUserController(services.UserService),
	}
}
