package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/services"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/utils"
)

type UserController struct {
	Manager services.ServiceInterface
}

func NewController(manager services.ServiceInterface) *UserController {
	return &UserController{
		Manager: manager,
	}
}

func (controller *UserController) CreateTask(c *gin.Context) {
	task := &models.Task{}
	err := c.BindJSON(task)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errCh := make(chan error)
	go controller.Manager.CreateTask(errCh, *task)
	err = <-errCh
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := &models.GeneralResponse{Success: true}
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
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := &models.GetTaskResponse{
		GeneralResponse: models.GeneralResponse{Success: true},
		Task:            <-taskCh,
	}
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

	errCh := make(chan error)
	go controller.Manager.UpdateTask(errCh, *task, taskId)
	err = <-errCh
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := &models.GeneralResponse{Success: true}
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
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := &models.GeneralResponse{Success: true}
	c.JSON(http.StatusOK, response)
}
