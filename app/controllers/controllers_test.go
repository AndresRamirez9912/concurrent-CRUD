package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/services"
)

type managerMock struct {
	Semaphore      chan bool
	UserRepository models.Repository
	TestResponse   error
}

func (managerMock *managerMock) CreateTask(errCh chan error, task models.Task) {
	errCh <- managerMock.TestResponse
}
func (managerMock *managerMock) GetTask(errCh chan error, taskId string, taskCh chan models.Task) {
	errCh <- managerMock.TestResponse
}
func (managerMock *managerMock) UpdateTask(errCh chan error, task models.Task, taskId string) {
	errCh <- managerMock.TestResponse
}
func (managerMock *managerMock) DeleteTask(errCh chan error, taskId string) {
	errCh <- managerMock.TestResponse
}

func (managerMock *managerMock) SignUpUser(errCh chan error, user *models.User) {
	errCh <- managerMock.TestResponse
}
func (managerMock *managerMock) LogInUser(errCh chan error, userCh chan *models.User, user string) {
	errCh <- managerMock.TestResponse
	userCh <- &models.User{Name: "test", Password: "Aaaaaeee3#V29928."}
}

var taskBody = &models.Task{
	Id:          "2",
	Title:       "Test Task",
	Description: "Test Description",
	State:       "pendiente",
}

var userBody = &models.User{
	Name:     "name test",
	Password: "Aaaaaeee3#V29928.",
}

func TestNewController(t *testing.T) {
	manager := services.NewManager(10, nil)
	controller := NewController(manager)

	_, ok := interface{}(controller.Manager).(services.ServiceInterface)
	if !ok {
		t.Error("Expected ServiceInterface, implemented", ok)
	}
}

func TestCreateTask(t *testing.T) {
	t.Run("Success result", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		controller.CreateTask(c)
		if w.Code != http.StatusOK {
			t.Error("Expected status code 200, got", w.Code)
		}
	})

	t.Run("Error parsing the body", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		errorBody := []byte(`{"errorTitle""Test Task","errorDescription":"Test Description","errorState":"Test State"}`)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(errorBody))
		controller.CreateTask(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400 bad request, got", w.Code)
		}
	})

	t.Run("Invalid task state", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		invalidTask := &models.Task{Id: "", Title: "", Description: "", State: "error"}
		invalidJson, _ := json.Marshal(invalidTask)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(invalidJson))
		controller.CreateTask(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400 bad request, got", w.Code)
		}
	})

	t.Run("Error Creating the Task", func(t *testing.T) {
		controller, c, w := InitMock(errors.New("test error"), taskBody)
		controller.CreateTask(c)
		if w.Code != http.StatusInternalServerError {
			t.Error("Expected status code 500 bad request, got", w.Code)
		}
	})
}

func TestGetTask(t *testing.T) {
	t.Run("Success result", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		go controller.GetTask(c)
		if w.Code != http.StatusOK {
			t.Error("Expected status code 200, got", w.Code)
		}
	})
}

func TestUpdateTask(t *testing.T) {

	t.Run("Success result", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		go controller.UpdateTask(c)
		if w.Code != http.StatusOK {
			t.Error("Expected status code 200, got", w.Code)
		}
	})

	t.Run("Error getting the task Id", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		url, _ := url.Parse("http://localhost:3000/tasks/ ")
		c.Request.URL = url
		controller.UpdateTask(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400 bad request, got", w.Code)
		}
	})

	t.Run("Invalid task state", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		invalidTask := &models.Task{Id: "", Title: "", Description: "", State: "error"}
		invalidJson, _ := json.Marshal(invalidTask)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(invalidJson))
		controller.UpdateTask(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400 bad request, got", w.Code)
		}
	})

	t.Run("Error parsing the body", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		errorBody := []byte(`{"errorTitle""Test Task","errorDescription":"Test Description","errorState":"Test State"}`)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(errorBody))
		controller.UpdateTask(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400 bad request, got", w.Code)
		}
	})
}

func TestDeleteTask(t *testing.T) {

	t.Run("Success result", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		go controller.DeleteTask(c)
		if w.Code != http.StatusOK {
			t.Error("Expected status code 200, got", w.Code)
		}
	})

	t.Run("Error getting the task Id", func(t *testing.T) {
		controller, c, w := InitMock(nil, taskBody)
		url, _ := url.Parse("http://localhost:3000/tasks/ ")
		c.Request.URL = url
		controller.DeleteTask(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400 bad request, got", w.Code)
		}
	})
}

func TestSignUp(t *testing.T) {

	t.Run("Success result", func(t *testing.T) {
		controller, c, w := InitMock(nil, userBody)
		controller.SignUp(c)
		if w.Code != http.StatusOK {
			t.Error("Expected status code 200, got", w.Code)
		}
	})

	t.Run("Error binding body", func(t *testing.T) {
		controller, c, w := InitMock(nil, userBody)
		c.Request.Body = nil
		controller.SignUp(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400, got", w.Code)
		}
	})

	t.Run("Error SignUp", func(t *testing.T) {
		controller, c, w := InitMock(errors.New("test error"), userBody)
		controller.SignUp(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400, got", w.Code)
		}
	})
}

func TestLogIn(t *testing.T) {
	t.Run("Success result", func(t *testing.T) {
		controller, c, w := InitMock(nil, userBody)
		controller.LogIn(c)
		if w.Code != http.StatusOK {
			t.Error("Expected status code 200, got", w.Code)
		}
	})

	t.Run("Error binding body", func(t *testing.T) {
		controller, c, w := InitMock(nil, userBody)
		c.Request.Body = nil
		controller.LogIn(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400, got", w.Code)
		}
	})

	t.Run("Error LogIn Operation", func(t *testing.T) {
		controller, c, w := InitMock(errors.New("test error"), userBody)
		controller.LogIn(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400, got", w.Code)
		}
	})

	t.Run("Error invalid credentials", func(t *testing.T) {
		errorBody := userBody
		errorBody.Password = "errorPassword"
		controller, c, w := InitMock(nil, errorBody)
		controller.LogIn(c)
		if w.Code != http.StatusBadRequest {
			t.Error("Expected status code 400, got", w.Code)
		}
	})
}

func InitMock(desiredError error, body any) (*UserController, *gin.Context, *httptest.ResponseRecorder) {
	managerMock := &managerMock{
		Semaphore:      make(chan bool, 2),
		UserRepository: nil,
		TestResponse:   desiredError,
	}

	jsonData, _ := json.Marshal(body)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controller := NewController(managerMock)

	url, _ := url.Parse("http://localhost:3000/tasks")
	c.Request = &http.Request{
		Header: make(http.Header),
		URL:    url,
		Method: http.MethodPost,
		Body:   io.NopCloser(bytes.NewBuffer(jsonData)),
	}

	return controller, c, w
}
