package services

import (
	"context"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
)

type UserService interface {
	GetUserByID(ctx context.Context, userUUID string) (*models.UserResponse, error)
	Register(ctx context.Context, req *models.RegisterRequest) (*models.UserResponse, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.UserResponse, error)
}

type UserSvc struct {
	repo repositories.UserRepository
}

func NewUserSvc() UserService {
	return &UserSvc{repo: repositories.NewUserRepository()}
}

func (us *UserSvc) Login(ctx context.Context, req *models.LoginRequest) (*models.UserResponse, error) {
	user, err := us.repo.LoginUser(req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return user, nil
}

func (us *UserSvc) Register(ctx context.Context, req *models.RegisterRequest) (*models.UserResponse, error) {
	user, err := us.repo.RegisterUser(req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return user, nil
}

func (us *UserSvc) GetUserByID(ctx context.Context, userUUID string) (*models.UserResponse, error) {
	user, err := us.repo.GetUserByID(userUUID)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return user, nil
}
