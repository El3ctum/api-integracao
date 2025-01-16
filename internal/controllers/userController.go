package controllers

import (
	// "github.com/gin-gonic/gin"
	services "api-integracao/internal/service"
)

type UserController struct {
	UserService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// func (uc *UserController) InsertDocumentoForUser() gin.HandlerFunc {
// 	return
// }
