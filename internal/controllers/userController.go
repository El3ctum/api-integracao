package controllers

import (
	// "github.com/gin-gonic/gin"
	"api-integracao/internal/models"
	services "api-integracao/internal/service"
	"fmt"

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

		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		err := uc.UserService.CreateUser(user.ID, &user)
		if err != nil {
			fmt.Println("Erro ao criar usuário")
			fmt.Println(err)
		}

		c.JSON(200, gin.H{"name": user.FirstName, "uuid": user.ID})
	}
}
