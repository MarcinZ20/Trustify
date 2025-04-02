package services

import (
	"github.com/MarcinZ20/Trustify/backend/api-service/api/models"
)

func PerformTextAnalysis(header string, content string, url string) (models.ResponseData, error) {
	responseData := models.ResponseData{
		Score:  78,
		Resutl: "True",
		Sources: []models.Source{
			{Url: "http://example.com"}, {Url: "https://something.com"},
		},
	}

	return responseData, nil
}
