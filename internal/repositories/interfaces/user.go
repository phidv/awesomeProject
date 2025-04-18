package interfaces

import (
	"context"
	"oms/internal/domain/models"
)

type User interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
}
