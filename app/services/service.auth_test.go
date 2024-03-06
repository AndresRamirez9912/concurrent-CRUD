package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/constants"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
)

func TestNewAuthService(t *testing.T) {
	auth := NewAuthService()
	if _, ok := interface{}(auth).(AuthInterface); !ok {
		t.Error("Expected AuthInterface, implemented", ok)
	}
}

func TestSendRequest(t *testing.T) {

	t.Run("Success path", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("token_value"))
		}))
		defer server.Close()
		err, token := SendRequest(server.URL, []byte("test_body"))
		if err != nil {
			t.Error("Error executing the sendRequest function", err)
		}

		if token != "token_value" {
			t.Errorf("Unexpected token: expected 'token_value', got %s ", token)
		}
	})

	t.Run("Fail create request", func(t *testing.T) {
		err, token := SendRequest("invalid_url", []byte("test_body"))
		if err == nil {
			t.Error("SendRequest should have returned an error")
		}
		if token != "" {
			t.Errorf("Unexpected token: expected empty, got %s", token)
		}
	})

	t.Run("Error reading body response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			hijacker, _ := w.(http.Hijacker)
			connection, _, _ := hijacker.Hijack()
			connection.Close()
		}))

		defer server.Close()

		err, token := SendRequest(server.URL, []byte("test_body"))
		if err == nil {
			t.Error("SendRequest should have returned an error")
		}
		if token != "" {
			t.Errorf("Unexpected token: expected empty, got %s", token)
		}
	})

	t.Run("Error sending the request", func(t *testing.T) {
		err, token := SendRequest("invalid_url", []byte("test_body"))
		if err == nil {
			t.Error("SendRequest should have returned an error")
		}
		if token != "" {
			t.Errorf("Unexpected token: expected empty, got %s", token)
		}
	})
}

func TestLogInAndSignUp(t *testing.T) {
	auth := NewAuthService()
	user := &models.User{
		Name:     "Test",
		Password: "test",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("token_value"))
	}))
	defer server.Close()

	t.Run("Success logIn or SignUp", func(t *testing.T) {
		token, err := auth.LogInAndSignUp(user, server.URL+"/login")
		if err != nil {
			t.Errorf("Error logging in: %v", err)
		}
		if token != "token_value" {
			t.Errorf("Unexpected token: expected 'token_value', got '%s'", token)
		}
	})

	t.Run("Error getting the user", func(t *testing.T) {
		_, err := auth.LogInAndSignUp(nil, "/login")
		if err == nil {
			t.Error("Expected Marshal error")
		}
	})

	t.Run("Error sending the request", func(t *testing.T) {
		_, err := auth.LogInAndSignUp(user, "/login")
		if err == nil {
			t.Error("Expected error sending the request")
		}
	})
}

func TestValidateUser(t *testing.T) {
	auth := NewAuthService()
	t.Run("Success validation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()
		err := auth.ValidateUser(server.URL, constants.TOKEN)
		if err != nil {
			t.Error("Expected success, the server return ok", err)
		}
	})

	t.Run("Token validation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}))
		defer server.Close()
		err := auth.ValidateUser(server.URL, constants.TOKEN)
		if err == nil {
			t.Error("Error expected, the server return not allowed", err)
		}
	})
}
