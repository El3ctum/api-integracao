package routes

import (
	"api-integracao/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HandleUsers(router *gin.RouterGroup, controllers controllers.ControllerInitializaer) {
	router.POST("/users", controllers.UserController.InsertDocumentForUser())
	router.GET("/users/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
