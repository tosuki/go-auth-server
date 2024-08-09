package users

import (
	"auth-server/internal/auth"
	"auth-server/internal/auth/model"
	"errors"

	"gorm.io/gorm"
)

type PostgresUserRepositoryImpl struct {
	Adapter *gorm.DB
}

func (repository *PostgresUserRepositoryImpl) Add(user *model.User) error {
	tx := repository.Adapter.Create(user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repository *PostgresUserRepositoryImpl) Has(email string) bool {
	_, err := repository.GetByEmail(email)

	if errors.Is(err, auth.ErrorInvalidEmail) {
		return false
	}

	if err != nil {
		panic(err)
	}

	return true
}

func (repository *PostgresUserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	var user model.User

	tx := repository.Adapter.First(&user, "email = ?", email)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, auth.ErrorInvalidEmail
	}

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
