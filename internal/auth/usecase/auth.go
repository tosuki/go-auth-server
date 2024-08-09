package usecase

import (
	"auth-server/internal/auth/repo/users"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("JWT_SECRET")

func SignUp(repository users.UserRepository, name, email string, password []byte) (string, error) {
    return "", nil
}

// string = jwt, error = if the request fails
func Authenticate(repository users.UserRepository, email string, password []byte) (string, error) {
	user, err := repository.GetByEmail(email)

	if err != nil {
		return "", fmt.Errorf("invalid-email")
	}

	isPasswordRight := bcrypt.CompareHashAndPassword([]byte(user.Password), password)

	if isPasswordRight != nil {
		return "", fmt.Errorf("invalid-password")
	}

	token, err := CreateToken(user.Name, user.Email, secretKey)

	if err != nil {
		return "", fmt.Errorf("token-err")
	}

	return token, nil
}
