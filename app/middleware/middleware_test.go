package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLimitGoroutines(t *testing.T) {
	// Prepare router
	router := gin.Default()
	router.Use(LimitGoroutines())
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Create request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// Create Recorder (request receiver)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Validate received data (in recorder)
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("Wrong status code: expected %v, got %v", http.StatusOK, status)
	}

	// Validate the redirected response
	body := recorder.Body.String()
	if body != "OK" {
		t.Errorf("Redirected Handler returned wrong body: expected %v, got %v", "OK", body)
	}
}
