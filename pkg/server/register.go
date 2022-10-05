package server

import (
	"github.com/gin-gonic/gin"
	"rest-api/pkg/repository/postgres"
)

func RegisterHTTPEndpoints(r *gin.Engine, ur *postgres.UserRepository) {
	h := NewHandler(ur)

	Endpoints := r.Group("/api/v1/")
	{
		Endpoints.GET("/user/:userId", h.GetUser)
		Endpoints.POST("/user", h.CreateUser)
		Endpoints.PUT("/user/:userId", h.UpdateUser)
		Endpoints.DELETE("/user/:userId", h.DeleteUser)
	}
}
