package controllers

import (
	"fmt"
	"net/http"

	"api-integracao/internal/auth"
	"api-integracao/internal/models"
	services "api-integracao/internal/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	AuthService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (ac *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		// Bind the incoming JSON to the `user` struct
		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Println("Error binding user JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"Response": "Invalid input"})
			return
		}

		// Check if the password is empty
		if user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Response": "Invalid credentials"})
			return
		}

		// Retrieve the user from the database using email
		userDb, err := ac.AuthService.GetUserByEmail(&user)
		if err != nil {
			fmt.Println("User not found:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"Response": "Invalid credentials"})
			return
		}

		// Compare the hashed password with the provided password
		err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))
		if err != nil {
			fmt.Println("Password mismatch")
			c.JSON(http.StatusUnauthorized, gin.H{"Response": "Invalid credentials"})
			return
		}

		// Prepare metadata for JWT
		userMetadata := models.UserMetadata{
			ID:          userDb.ID,
			Name:        fmt.Sprintf("%s %s", userDb.FirstName, userDb.LastName),
			Companies:   userDb.Companies,
			Departments: userDb.Departments,
			Roles:       userDb.Roles,
			Permissions: userDb.Permissions,
		}

		// Generate a JWT token
		token, err := auth.GenerateJwtToken(userMetadata)
		if err != nil {
			fmt.Println("Error generating JWT token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Response": "Authentication failed"})
			return
		}

		// Respond with success
		c.JSON(http.StatusOK, gin.H{
			"Response": fmt.Sprintf("User authenticated successfully: %s", userDb.FirstName),
			"Token":    token,
		})
	}
}
