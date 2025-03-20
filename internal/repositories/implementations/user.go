package implementations

import (
	"context"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"oms/internal/domain/models"
	"oms/internal/infrastructure/database"
	"oms/internal/repositories/interfaces"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.User {
	return &UserRepository{db: db}
}

func (repo UserRepository) StandardQuery() *gorm.DB {
	return repo.db
}

func (UserRepository) Create(ctx context.Context, user *models.User) error {
	return database.Transaction(ctx, func(tx *gorm.DB) error {
		return database.PostgresqlDB(ctx).Create(user).Error
	})
}

func (repo UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := repo.StandardQuery().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	user := models.User{ID: uuid.Must(uuid.FromString(id))}
	if err := repo.StandardQuery().First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
