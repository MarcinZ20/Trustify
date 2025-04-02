package routes

import (
	"github.com/MarcinZ20/Trustify/backend/api-service/api/handlers"
	"github.com/gin-gonic/gin"
)

func ServerRoutes(router *gin.Engine) {
	{
		v1 := router.Group("/api/v1")
		v1.POST("/search", handlers.HandleSearch)
		v1.GET("/health", handlers.HealthCheck)
	}
}
