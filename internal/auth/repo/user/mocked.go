package user

import (
	"go-auth-server/internal/auth"
	"go-auth-server/internal/auth/models"
)

type MockedUserRepositoryImpl struct {
	Adapter map[string]models.User
}

func (repository *MockedUserRepositoryImpl) Add(user *models.User) error {
	_, ok := repository.Adapter[user.Email]

	if ok {
		return auth.ErrDuplicatedUser
	}

	repository.Adapter[user.Email] = *user

	return nil
}

func (repository *MockedUserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	value, ok := repository.Adapter[email]

	if !ok {
		return nil, auth.ErrInvalidEmail
	}

	return &value, nil
}

func (repository *MockedUserRepositoryImpl) Delete(email string) error {
	return nil
}

func NewMockedUserRepository() UserRepository {
	return &MockedUserRepositoryImpl{
		Adapter: make(map[string]models.User),
	}
}
