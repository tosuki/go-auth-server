package users

import (
	"auth-server/internal/auth/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user *model.User) error
	Has(email string) bool
	GetByEmail(email string) (*model.User, error)
}

func NewUserRepository(adapter *gorm.DB) UserRepository {
	adapter.AutoMigrate(&model.User{})

	return &PostgresUserRepositoryImpl{
		Adapter: adapter,
	}
}
