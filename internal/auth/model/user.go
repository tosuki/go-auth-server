package model

type User struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt int64
	UpdatedAt int64
}

func NewUser(id, name, email, password string, createdAt, updatedAt int64) *User {
	user := User{
		Id:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return &user
}
