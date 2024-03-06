package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/constants"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/utils"
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

func TestValidateUserMiddleware(t *testing.T) {

	t.Run("Valid token", func(t *testing.T) {
		router := gin.New()
		router.Use(ValidateUser(true))
		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		validToken := utils.CreateJWT(&models.User{Name: "test", Password: "Aa@123"})
		req.AddCookie(&http.Cookie{Name: constants.TOKEN, Value: validToken})

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("Doesn't have token", func(t *testing.T) {
		router := gin.New()
		router.Use(ValidateUser(true))

		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		router := gin.New()
		router.Use(ValidateUser(true))
		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		validToken := utils.CreateJWT(&models.User{Name: "test", Password: "Aa@123"})
		req.AddCookie(&http.Cookie{Name: constants.TOKEN, Value: validToken + "aaaa"})

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		router.ServeHTTP(w, req)
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d but got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})

	t.Run("Inactive auth validation", func(t *testing.T) {
		router := gin.New()
		router.Use(ValidateUser(false))
		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
		}
	})
}
