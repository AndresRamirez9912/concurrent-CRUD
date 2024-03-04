package utils

import (
	"errors"
	"net/http"
	"strings"

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

func ValidateTaskState(task *models.Task) error {
	validStatus := map[string]bool{
		"pendiente":   true,
		"en progreso": true,
		"completada":  true,
	}
	result := validStatus[strings.ToLower(task.State)]
	if !result {
		return errors.New("Task state, not allowed")
	}
	return nil
}
