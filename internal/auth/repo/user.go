package repo

import "auth-server/internal/auth/model"

type UserRepository interface {
	Save(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}

type UserRepositoryImpl struct {
	Adapter interface{}
}

func (repository *UserRepositoryImpl) Save(user *model.User) error {
	return nil
}

func (repository *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
    return nil, nil
}
