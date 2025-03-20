package repositories

import (
	"gorm.io/gorm"
	"oms/internal/repositories/implementations"
	"oms/internal/repositories/interfaces"
)

type Repositories struct {
	UserRepository interfaces.User
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: implementations.NewUserRepository(db),
	}
}
