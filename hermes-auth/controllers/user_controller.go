package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/xu3cl40122/hermes/hermes-auth/models"
	"github.com/xu3cl40122/hermes/hermes-auth/services"
)
type UserController struct {
	userService services.UserService
}
func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (userController *UserController) CreateUser(ctx *gin.Context) {
	var req = models.CreateUserInput{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := userController.userService.CreateUser(ctx,  &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Done()
	return
}
// RegisterUser 用戶註冊
func (userController *UserController) Login (ctx *gin.Context) {
	var req = models.LoginInput{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := userController.userService.Login(ctx,  &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
	return
}
