package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, s *Server) {
	r.GET("/login", s.Login)
	r.GET("/signup", s.Signup)
}

func (s *Server) Login(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello Davi"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) Signup(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello Davi"

	c.JSON(http.StatusOK, resp)
}