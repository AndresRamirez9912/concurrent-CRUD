package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/controllers"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
)

func main() {
	db := models.NewSqlConnection()
	defer db.CloseConnection()

	controller := controllers.NewController(db)

	router := gin.Default()
	router.POST("/tasks", controller.CreateTask)
	router.GET("/tasks/:id", controller.GetTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.DELETE("/tasks/:id", controller.DeleteTask)

	s := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Error trying to starting the server", err)
	}
}
