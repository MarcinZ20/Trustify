package services

import (
	"fmt"

	"github.com/MarcinZ20/Trustify/backend/api-service/api/models"
	"github.com/MarcinZ20/Trustify/backend/api-service/config/client"
)

func PerformTextAnalysis(header string, content string, url string) (*models.ResponseData, error) {
	apiClient := client.GetClientConfig()

	var respData models.ResponseData

	payload := map[string]string{
		"headline": header,
		"content":  content,
		"url":      url,
	}

	response, err := apiClient.Client.POST("/search").Body().AsJSON(payload).Send()
	if err != nil {
		return &respData, fmt.Errorf("bad request: %s", err)
	}

	if err := response.Body().AsJSON(&respData); err != nil {
		return &respData, fmt.Errorf("parsing error: %s", err)
	}

	return &respData, nil
}
