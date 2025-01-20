package controllers

import (
	"fmt"
	"net/http"

	"api-integracao/internal/models"
	services "api-integracao/internal/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

		// Check if the password is empty
		if user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Response": "Invalid inputs"})
			return
		}

		// Hash the password to save on the database
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Password hashing has failed:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Response": "User could not be created now, try again later"})
			return
		}

		user.Password = string(hashPassword)

		err = uc.UserService.CreateUser(*user.ID, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Response": "user not created, try again later!"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"Response": fmt.Sprintf("User created with sucess: %s", user.FirstName)})
	}
}

func (uc *UserController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		userDB, err := uc.UserService.GetUserById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Response": "user not found, check if the user exists!"})
			return
		}

		userResponse := models.UserMetadata{
			ID:          *userDB.ID,
			Name:        fmt.Sprintf("%s %s", userDB.FirstName, userDB.LastName),
			Companies:   userDB.Companies,
			Departments: userDB.Departments,
			Roles:       userDB.Roles,
			Permissions: userDB.Permissions,
		}

		c.JSON(http.StatusOK, gin.H{"Response": userResponse})
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

		err := uc.UserService.CreateUser(*user.ID, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Response": "user not created, try again later!"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"Response": fmt.Sprintf("User created with sucess: %s", user.FirstName)})
	}
}
