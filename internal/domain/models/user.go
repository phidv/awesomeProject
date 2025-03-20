package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"oms/pkg/utils"
)

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
}

// BeforeCreate là hook để hash password trước khi lưu vào database
func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
