package user

import (
	"errors"
	"go-auth-server/internal/auth"
	"go-auth-server/internal/auth/models"

	"gorm.io/gorm"
)

type GormUserRepositoryImpl struct {
	Adapter *gorm.DB
}

func (repository *GormUserRepositoryImpl) Add(user *models.User) error {
	tx := repository.Adapter.Create(user)

	if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
		return auth.ErrDuplicatedUser
	}

	return tx.Error
}

func (repository *GormUserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User

	tx := repository.Adapter.First(&user, "email = ?", email)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, auth.ErrInvalidEmail
	}

	if tx.Error != nil {
	}

	return &user, nil
}

func (repository *GormUserRepositoryImpl) Delete(email string) error {
	return nil
}

func NewUserRepository(adapter *gorm.DB) UserRepository {
	return &GormUserRepositoryImpl{
		Adapter: adapter,
	}
}
