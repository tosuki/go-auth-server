package repo

import (
	"auth-server/internal/auth/model"
	"fmt"
)

type UserRepository interface {
	Add(user *model.User) error
	Has(email string) bool
	GetByEmail(email string) (*model.User, error)
}

type FakeUserRepositoryImpl struct {
	Adapter map[string]model.User
}

func (repository *FakeUserRepositoryImpl) Has(email string) bool {
	_, err := repository.Adapter[email]

	return !err
}

func (repository *FakeUserRepositoryImpl) Add(user *model.User) error {
	if repository.Has(user.Email) {
		return fmt.Errorf("unique-signature-error")
	}

	repository.Adapter[user.Email] = *user

	return nil
}

func (repository *FakeUserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	user, err := repository.Adapter[email]

	if err {
		return nil, fmt.Errorf("invalid-email")
	}

	return &user, nil
}

func NewUserRepository() UserRepository {
	return &FakeUserRepositoryImpl{
		Adapter: make(map[string]model.User),
	}
}
