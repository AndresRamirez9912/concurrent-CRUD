package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/utils"
)

type UserController struct {
	UserRepository models.Repository
}

func NewController(repo models.Repository) *UserController {
	return &UserController{
		UserRepository: repo,
	}
}

func (controller *UserController) CreateTask(c *gin.Context) {
	task := &models.Task{}
	err := c.BindJSON(task)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
	}

	go func() {
		err = controller.UserRepository.CreateTask(*task)
		if err != nil {
			errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, errorResponse)
		}
	}()

	response := &models.GeneralResponse{Success: true}
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) GetTask(c *gin.Context) {
	taskId := c.Param("id")
	utils.ValidateTestId(taskId, c)

	task, err := controller.UserRepository.GetTask(taskId)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
	}

	response := &models.GetTaskResponse{
		GeneralResponse: models.GeneralResponse{Success: true},
		Task:            *task,
	}
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) UpdateTask(c *gin.Context) {
	taskId := c.Query("id")
	utils.ValidateTestId(taskId, c)

	task := &models.Task{}
	err := c.BindJSON(task)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
	}

	err = controller.UserRepository.UpdateTask(*task, taskId)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
	}

	response := &models.GeneralResponse{Success: true}
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	utils.ValidateTestId(taskId, c)

	err := controller.UserRepository.DeleteTask(taskId)
	if err != nil {
		errorResponse := utils.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
	}

	response := &models.GeneralResponse{Success: true}
	c.JSON(http.StatusOK, response)
}
