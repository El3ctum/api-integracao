package routes

import (
	"api-integracao/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HandleAuth(router *gin.RouterGroup, controllers controllers.ControllerInitializer) {
	router.POST("/login", controllers.AuthController.Login())
	router.POST("/register", controllers.AuthController.Login())
}
