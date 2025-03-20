package handlers

import "oms/internal/services"

type AppHandlers struct {
	AuthHandler *AuthHandler
	UserHandler *UserHandler
}

func NewAppHandlers(services *services.AppServices) *AppHandlers {
	return &AppHandlers{
		AuthHandler: NewAuthHandler(*services.AuthService),
		UserHandler: NewUserHandler(*services.UserService),
	}
}
