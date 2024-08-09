package usecase

import (
	"auth-server/internal/auth"
	"auth-server/internal/auth/model"
	"auth-server/internal/auth/repo/users"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("JWT_SECRET")

type AuthUsecase struct {
	Repository users.UserRepository
}

func (usecase *AuthUsecase) SignUp(name, email string, password []byte) (string, error) {
	_, queryErr := usecase.Repository.GetByEmail(email)

	if !errors.Is(queryErr, auth.ErrorInvalidEmail) {
		return "", auth.ErrorEmailTaken
	}

	if queryErr != nil {
		return "", auth.ErrorUnsupported
	}

	userUUID, uuidErr := uuid.NewV7()

	if uuidErr != nil {
		return "", uuidErr
	}

	hashedPassword, encryptErr := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if encryptErr != nil {
		return "", encryptErr
	}

	now := time.Now().Unix()

	usecase.Repository.Add(model.NewUser(userUUID.String(), name, email, string(hashedPassword), now, now))

	token, tokenErr := CreateToken(name, email, secretKey)

	if tokenErr != nil {
		return "", auth.ErrorTokenCreationError
	}

	return token, nil
}

// string = jwt, error = if the request fails
func (usecase *AuthUsecase) Authenticate(email string, password []byte) (string, error) {
	user, err := usecase.Repository.GetByEmail(email)

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
