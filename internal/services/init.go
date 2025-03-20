package services

import "oms/internal/repositories"

type AppServices struct {
	AuthService *AuthService
	UserService *UserService
}

func NewAppServices(repo *repositories.Repositories) *AppServices {
	return &AppServices{
		AuthService: NewAuthService(*repo),
		UserService: NewUserService(*repo),
	}
}
