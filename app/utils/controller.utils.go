package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/constants"
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

func CreateJWT(user *models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["sub"] = user.Name

	secretKey := []byte(constants.JWT_SECRET)
	tokenString, _ := token.SignedString(secretKey)
	return tokenString
}

func DecriptJWT(JWT string) error {
	secretKey := []byte(constants.JWT_SECRET)
	token, err := jwt.Parse(JWT, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("Invalid token, please logIn")
	}
	return nil
}
