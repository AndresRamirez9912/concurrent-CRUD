package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/constants"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/services"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/utils"
)

type UserController struct {
	Manager services.ServiceInterface
	Auth    services.AuthInterface
}

func NewController(manager services.ServiceInterface, auth services.AuthInterface) *UserController {
	return &UserController{
		Manager: manager,
		Auth:    auth,
	}
}

func (controller *UserController) CreateTask(c *gin.Context) {
	task := &models.Task{}
	err := c.BindJSON(task)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if err = utils.ValidateTaskState(task); err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errCh := make(chan error)
	go controller.Manager.CreateTask(errCh, *task)
	err = <-errCh
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := &models.GeneralResponse{Success: true}
	c.Header("Content-Security-Policy", "default-src 'self'")
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) GetTask(c *gin.Context) {
	taskId := c.Param("id")
	utils.ValidateTestId(taskId, c)

	errCh := make(chan error)
	taskCh := make(chan models.Task)
	go controller.Manager.GetTask(errCh, taskId, taskCh)
	err := <-errCh
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := &models.GetTaskResponse{
		GeneralResponse: models.GeneralResponse{Success: true},
		Task:            <-taskCh,
	}
	c.Header("Content-Security-Policy", "default-src 'self'")
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	utils.ValidateTestId(taskId, c)
	task := &models.Task{}
	err := c.BindJSON(task)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if err = utils.ValidateTaskState(task); err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errCh := make(chan error)
	go controller.Manager.UpdateTask(errCh, *task, taskId)
	err = <-errCh
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := &models.GeneralResponse{Success: true}
	c.Header("Content-Security-Policy", "default-src 'self'")
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	utils.ValidateTestId(taskId, c)

	errCh := make(chan error)
	go controller.Manager.DeleteTask(errCh, taskId)
	err := <-errCh
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := &models.GeneralResponse{Success: true}
	c.Header("Content-Security-Policy", "default-src 'self'")
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) SignUp(c *gin.Context) {
	user := &models.User{}
	err := c.BindJSON(user)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	token, err := controller.Auth.LogInAndSignUp(user, constants.SIGNUP_URL)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	c.SetCookie(constants.TOKEN, token, 3600, "", "", true, true)
	response := &models.GeneralResponse{Success: true}
	c.Header("Content-Security-Policy", "default-src 'self'")
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) LogIn(c *gin.Context) {
	user := &models.User{}
	err := c.BindJSON(user)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	token, err := controller.Auth.LogInAndSignUp(user, constants.LOGIN_URL)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	c.SetCookie(constants.TOKEN, token, 3600, "/", "", true, true)
	response := &models.GeneralResponse{Success: true}
	c.Header("Content-Security-Policy", "default-src 'self'")
	c.JSON(http.StatusOK, response)
}
