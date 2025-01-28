package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xu3cl40122/hermes.git/hermes-auth/middleware"
	"github.com/xu3cl40122/hermes/hermes-auth/controllers"
	"github.com/xu3cl40122/hermes/hermes-auth/infra"
	"github.com/xu3cl40122/hermes/hermes-auth/services"
)

func main() {
	infra.ConnectDB()
	userRepo := infra.NewMongoUserRepository(infra.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	// Apply middleware
	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/profile", userController.GetProfile)
	}

	r.POST("/register", userController.CreateUser)
	r.POST("/login", userController.Login)

	r.Run(":8080")
}
