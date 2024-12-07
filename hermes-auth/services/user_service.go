package services

import (
	"context"
	"time"
	"github.com/xu3cl40122/hermes/hermes-auth/infra"
	"github.com/xu3cl40122/hermes/hermes-auth/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, input * models.CreateUserInput) (*models.User, error)
}
type UserServiceImpl struct {
	repo infra.UserRepository
}
func NewUserService(repo infra.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, input *models.CreateUserInput) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	newUser := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Nickname: input.Nickname,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = service.repo.Create(ctx, &newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}