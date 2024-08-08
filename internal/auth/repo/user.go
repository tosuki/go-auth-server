package repo

import (
    "auth-server/internal/auth/model"
    "time"
)

type UserRepository interface {
	Save(user *model.User) (model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type UserRepositoryImpl struct {
	Adapter map[string]model.User
}

func NewUserRepository() *UserRepository {
    return &UserRepositoryImpl{
        Adapter: make(map[string]model.User)
    }
}

func (repository *UserRepositoryImpl) Has(email string) bool {
    
}

func (repository *UserRepositoryImpl) Save(id, name, email, password string) (model.User, error) {
    now := time.Now().Unix()

    user := model.User{
        Id: id,
        Email: email,
        Password: password,
        CreatedAt: now,
    }

    repository.Adapter[email] = user

    return nil, nil
}

func (repository *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
    return nil, nil
}
