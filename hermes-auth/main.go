package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xu3cl40122/hermes/hermes-auth/controllers"
	"github.com/xu3cl40122/hermes/hermes-auth/infra"
)

func main() {
	// 初始化 MongoDB 連接
	infra.ConnectDB()
	// 初始化 UserRepository
	userRepo := infra.NewMongoUserRepository(infra.DB)

	// 初始化 UserController
	userController := controllers.NewUserController(userRepo)

	// 創建 Gin 路由
	r := gin.Default()

	// 設置路由
	r.POST("/user", userController.CreateUser)

	// 啟動服務
	r.Run(":8080")
}
