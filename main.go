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

	// Configuration
	manager := services.NewManager(50, db)
	auth := services.NewAuthService()
	controller := controllers.NewController(manager, auth)

	// Router
	router := gin.Default()

	// Endpoints
	router.POST("/tasks", middleware.ValidateUser(true, auth), controller.CreateTask)
	router.GET("/tasks/:id", middleware.ValidateUser(true, auth), controller.GetTask)
	router.PUT("/tasks/:id", middleware.ValidateUser(true, auth), controller.UpdateTask)
	router.DELETE("/tasks/:id", middleware.ValidateUser(true, auth), controller.DeleteTask)

	router.Use(middleware.LimitGoroutines())

	router.POST("/signUp", controller.SignUp)
	router.POST("/logIn", controller.LogIn)

	s := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Error trying to starting the server", err)
	}
}
