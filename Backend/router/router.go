package router

import (
	"hope-backend-gin/router/health"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	root := r.Group("/")
	health.RegisterRoutes(root)
	
	// Future: v1 := r.Group("/api/v1")
	// users.RegisterRoutes(v1)

	return r
}
