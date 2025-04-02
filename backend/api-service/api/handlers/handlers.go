package handlers

import (
	"net/http"

	"github.com/MarcinZ20/Trustify/backend/api-service/api/models"
	"github.com/MarcinZ20/Trustify/backend/api-service/internal/services"
	"github.com/gin-gonic/gin"
)

func HandleSearch(ctx *gin.Context) {
	var requestData models.RequestData

	if err := ctx.BindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
	}

	result, err := services.PerformTextAnalysis(requestData.Headline, requestData.Content, requestData.Url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to perform the analysis",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "HEALTHY",
	})
}
