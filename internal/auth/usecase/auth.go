package usecase

import (
	"auth-server/internal/auth"
	"auth-server/internal/auth/repo/users"
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
		return "", auth.ErrorInvalidEmail
	}

	isPasswordRight := bcrypt.CompareHashAndPassword([]byte(user.Password), password)

	if isPasswordRight != nil {
		return "", auth.ErrorInvalidPassword
	}

	token, err := CreateToken(user.Name, user.Email, secretKey)

	if err != nil {
		return "", auth.ErrorTokenCreationError
	}

	return token, nil
}
