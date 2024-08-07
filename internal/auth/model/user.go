package model

type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt int
	UpdatedAt int
}

func NewUser(id, name, email, password string, createdAt, updatedAt int) (*User, error) {
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
