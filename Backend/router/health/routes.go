package health

import (
	"hope-backend-gin/router/response"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", func(c *gin.Context) {
		response.OK(c, gin.H{"status": "ok"})
	})
}