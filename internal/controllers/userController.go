package controllers

import (
	// "github.com/gin-gonic/gin"
	"api-integracao/internal/models"
	services "api-integracao/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) InsertDocumentForUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := uc.UserService.CreateUser("teste", &models.User{
			ID:          "Teste",
			FirstName:   "Davi",
			LastName:    "vieira",
			Email:       "davilealmarcal198@gmail.com",
			Password:    "zezinho123",
			Companies:   []string{"Company1", "Company2"},
			Departments: []string{"Processos", "RPA"},
			Roles:       []string{"Admin", "Analista"},
			Permissions: []string{"read", "delete"},
		})
		if err != nil {
			fmt.Println("Erro ao criar usuário")
		}
		c.JSON(http.StatusCreated, gin.H{
			"User": "User created with sucess",
		})
	}
}
