package users

import (
	"auth-server/internal/auth/model"

	"gorm.io/gorm"
)

type PostgresUserRepositoryImpl struct {
	Adapter *gorm.DB
}

func (repository *PostgresUserRepositoryImpl) Add(user *model.User) error {
	return nil
}

func (repository *PostgresUserRepositoryImpl) Has(email string) bool {
	return false
}

func (repository *PostgresUserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	return nil, nil
}
