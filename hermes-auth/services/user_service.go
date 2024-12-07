package services

import (
	"context"
	"errors"
	"time"
	"github.com/xu3cl40122/hermes/hermes-auth/infra"
	"github.com/xu3cl40122/hermes/hermes-auth/models"
	auth "github.com/xu3cl40122/hermes/hermes-auth/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, input * models.CreateUserInput) (*models.User, error)
	Login(ctx context.Context, input * models.LoginInput) (string, error)
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

func (service *UserServiceImpl) Login(ctx context.Context, input *models.LoginInput) (string, error) {
	user, err := service.repo.Get(ctx, input.Email)
	if err != nil {
			return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
			return "", errors.New("invalid password")
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
			return "", err
	}

	return token, nil
}