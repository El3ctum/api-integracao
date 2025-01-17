package routes

import (
	"api-integracao/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HandleUsers(router *gin.RouterGroup, controllers controllers.ControllerInitializaer) {
	router.GET("/users/:id", controllers.UserController.InsertDocumentForUser())
}
