package controllers

import service "api-integracao/internal/service"

type ControllerInitializer struct {
	AuthController *AuthController
	UserController *UserController
}

func InitControllers(services service.ServiceInitializaer) ControllerInitializer {
	return ControllerInitializer{
		AuthController: NewAuthController(services.AuthService),
		UserController: NewUserController(services.UserService),
	}
}
