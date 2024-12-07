package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xu3cl40122/hermes/hermes-auth/controllers"
	"github.com/xu3cl40122/hermes/hermes-auth/services"
	"github.com/xu3cl40122/hermes/hermes-auth/infra"
)

func main() {
	infra.ConnectDB()
	userRepo := infra.NewMongoUserRepository(infra.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	r.POST("/register", userController.CreateUser)
	r.POST("/login", userController.Login)

	r.Run(":8080")
}
