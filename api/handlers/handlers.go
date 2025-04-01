package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSearch(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/search endpoint",
	})
}

func HealthCheck(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "HEALTHY",
  })
}
