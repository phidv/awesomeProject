package services

import (
	"context"
	"errors"
	"oms/internal/api/dto/request"
	"oms/internal/domain/models"
	"oms/internal/repositories"
	"oms/internal/repositories/interfaces"
	"oms/pkg/utils"
)

type AuthService struct {
	UserRepo interfaces.User
}

func NewAuthService(repo repositories.Repositories) *AuthService {
	return &AuthService{UserRepo: repo.UserRepository}
}

func (s *AuthService) Register(ctx context.Context, req request.RegisterRequest) (*models.User, error) {
	existingUser, _ := s.UserRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Authenticate(ctx context.Context, req request.LoginRequest) (token string, err error) {
	user, err := s.UserRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	match := utils.ComparePassword(user.Password, req.Password)
	if !match {
		return "", errors.New("invalid email or password")
	}

	return utils.GenerateToken(user.ID.String(), user.Email)
}
