package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateErrorResponse(t *testing.T) {
	errorResponse := CreateErrorResponse(1234, "testing")
	if errorResponse.ErrorCode != 1234 {
		t.Errorf("Wrong errorCode. Expected: %d, obtained: %d", 1234, errorResponse.ErrorCode)
	}

	if errorResponse.Success == true {
		t.Errorf("Wrong success. Expected: %t, obtained: %t", false, errorResponse.Success)
	}

	if errorResponse.ErrorMessage != "testing" {
		t.Errorf("Wrong success. Expected: %s, obtained: %s", "testing", errorResponse.ErrorMessage)
	}
}

func TestValidateTestId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	t.Run("Id is not empty", func(t *testing.T) {
		ValidateTestId("1", c)
		if w.Code == http.StatusBadRequest {
			t.Errorf("Wrong statusCode. Expected: %d, obtained: %d", 200, w.Code)
		}
	})

	t.Run("Id is empty", func(t *testing.T) {
		ValidateTestId("", c)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Wrong statusCode. Expected: %d, obtained: %d", http.StatusBadRequest, w.Code)
		}
	})
}
