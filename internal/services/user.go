package services

import (
	"context"
	"oms/internal/api/dto/request"
	"oms/internal/domain/models"
	"oms/internal/repositories"
	"oms/internal/repositories/interfaces"
)

type UserService struct {
	UserRepo interfaces.User
}

func NewUserService(repo repositories.Repositories) *UserService {
	return &UserService{UserRepo: repo.UserRepository}
}

func (s *UserService) Get(ctx context.Context, req request.GetUserRequest) (*models.User, error) {
	return s.UserRepo.FindByID(ctx, req.Id)
}

func (s *UserService) List(ctx context.Context, req request.GetUserRequest) (*models.User, error) {
	return s.UserRepo.FindByID(ctx, req.Id)
}
