package routes

import (
	"api-integracao/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HandleUsers(router *gin.RouterGroup, controllers controllers.ControllerInitializer) {
	router.POST("/users", controllers.UserController.InsertDocumentForUser())
	router.GET("/users/:id", controllers.UserController.GetUserById())
}
