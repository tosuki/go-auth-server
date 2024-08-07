package model

type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
}

func NewUser(id, name, email, password string, createdAt, updatedAt int64) (*User, error) {
	user := User{
		Id:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return &user, nil
}
