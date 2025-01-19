package controllers

import (
	"fmt"
	"net/http"

	"api-integracao/internal/models"
	services "api-integracao/internal/service"

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
			c.JSON(http.StatusBadRequest, gin.H{"Response": err.Error()})
			return
		}

		err := uc.UserService.CreateUser(user.ID, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Response": "user not created, try again later!"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"Response": fmt.Sprintf("User created with sucess: %s", user.FirstName)})
	}
}

func (uc *UserController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"Response": err.Error()})
			return
		}

		err := uc.UserService.CreateUser(user.ID, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Response": "user not created, try again later!"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"Response": fmt.Sprintf("User created with sucess: %s", user.FirstName)})
	}
}

func (uc *UserController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"Response": err.Error()})
			return
		}

		err := uc.UserService.CreateUser(user.ID, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Response": "user not created, try again later!"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"Response": fmt.Sprintf("User created with sucess: %s", user.FirstName)})
	}
}
