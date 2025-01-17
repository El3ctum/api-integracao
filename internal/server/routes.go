package server

import (
	"api-integracao/internal/helpers"
	"api-integracao/internal/routes"
	"net/http"

	// services "api-integracao/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	controllers := helpers.InitControllers(s.db.GetScope())

	v1 := r.Group("/v1")
	routes.HandleUsers(v1, controllers)

	r.GET("/", s.HelloWorldHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	// resp := make(map[string]string)
	// resp["message"] = "Hello World"
	resp := s.db.Health()

	c.JSON(http.StatusOK, resp)
}
