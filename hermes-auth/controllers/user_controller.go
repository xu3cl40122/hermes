package controllers

import (
	"net/http"
	"github.com/xu3cl40122/hermes/hermes-auth/infra"
	"github.com/xu3cl40122/hermes/hermes-auth/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
type UserController struct {
	repo infra.UserRepository
}
func NewUserController(repo infra.UserRepository) *UserController {
	return &UserController{repo: repo}
}

// RegisterUser 用戶註冊
func (uc *UserController) CreateUser(c *gin.Context) {
	// 定義請求結構體
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Nickname string `json:"nickname" binding:"required"`
	}

	// 解析請求
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密碼
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// 創建新用戶
	newUser := models.User{
		Email:    request.Email,
		Password: string(hashedPassword),
		Nickname: request.Nickname,
	}

	err = uc.repo.Create(c, &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User created"})
	return
}
