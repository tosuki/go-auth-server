package models

import (
	"go-auth-server/internal/auth"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string

	CreatedAt int64
	UpdatedAt int64
}

func NewUser(name, email, password string) (*User, error) {
	uuid, uuidErr := uuid.NewV7()

	if uuidErr != nil {
		return nil, auth.ErrFailedToCreateUUID
	}

	now := time.Now().Unix()

	return &User{
		Id:        uuid.String(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
