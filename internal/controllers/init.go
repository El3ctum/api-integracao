package controllers

import service "api-integracao/internal/service"

type ControllerInitializer struct {
	UserController *UserController
}

func InitControllers(services service.ServiceInitializaer) ControllerInitializer {
	return ControllerInitializer{
		UserController: NewUserController(services.UserService),
	}
}
