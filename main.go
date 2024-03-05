package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/controllers"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/middleware"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/services"
)

func main() {
	db := models.NewSqlConnection()
	defer func() {
		err := db.CloseConnection()
		if err != nil {
			log.Fatal("Error closing the connection", err)
		}
	}()

	manager := services.NewManager(50, db)
	controller := controllers.NewController(manager)
	router := gin.Default()
	router.Use(middleware.LimitGoroutines())
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
