package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/constants"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
)

type AuthInterface interface {
	LogInAndSignUp(*models.User, string) (string, error)
	ValidateUser(string, string) error
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (auth *AuthService) LogInAndSignUp(user *models.User, url string) (string, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	err, token := SendRequest(url, jsonData)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (auth *AuthService) ValidateUser(urlLambda, token string) error {
	url, _ := url.Parse(urlLambda)
	query := url.Query()
	query.Add(constants.TOKEN, token)
	url.RawQuery = query.Encode()

	err, _ := SendRequest(url.String(), nil)
	if err != nil {
		return err
	}
	return nil
}

func SendRequest(urlReceived string, body []byte) (error, string) {
	url, _ := url.Parse(urlReceived)
	operation := strings.Split(url.Path, "/")
	messages := map[string]string{"signUp": "User already exists", "logIn": "Invalid credentiald"}
	req, err := http.NewRequest(http.MethodPost, urlReceived, bytes.NewBuffer(body))
	if err != nil {
		return err, ""
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(messages[operation[len(operation)-1]]), ""
	}

	token, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}
	return nil, string(token)
}
