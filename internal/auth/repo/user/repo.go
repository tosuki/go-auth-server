package user

import "go-auth-server/internal/auth/models"

type UserRepository interface {
	Add(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	Delete(email string) error
}
