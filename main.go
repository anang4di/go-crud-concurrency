package main

import (
	"go-crud-concurrency/handler"
	"go-crud-concurrency/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	userRepository := user.NewRepository()
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	api := router.Group("/api/v1")

	api.GET("/users", userHandler.GetAllUsers)
	api.POST("/users", userHandler.RegisterUser)

	router.Run()

}
