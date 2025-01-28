package controllers

import (
	"log"
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

	_, err := userController.userService.CreateUser(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Done()
}

func (userController *UserController) Login(ctx *gin.Context) {
	var req = models.LoginInput{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := userController.userService.Login(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserController) GetProfile(ctx *gin.Context) {
	authUser, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "session not found"})
		return
	}
	
	c, ok := authUser.(*models.UserClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found 888"})
		return
	}
	user, err := u.userService.GetById(ctx, c.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
