package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
)

func CreateErrorResponse(errorCode int, errorMessage string) *models.GeneralResponse {
	return &models.GeneralResponse{
		Success:      false,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}

func ValidateTestId(taskId string, c *gin.Context) {
	if taskId == "" {
		err := errors.New("Invalid task ID")
		errorResponse := CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
	}
}
