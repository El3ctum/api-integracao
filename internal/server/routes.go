package server

import (
	"api-integracao/internal/helpers"
	"api-integracao/internal/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	controllers := helpers.InitControllers(s.db.GetScope())

	v1 := r.Group("/v1")
	routes.HandleAuth(v1, controllers)
	routes.HandleUsers(v1, controllers)

	return r
}
