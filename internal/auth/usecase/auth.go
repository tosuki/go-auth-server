package usecase

import (
	"errors"
	"go-auth-server/internal/auth"
	"go-auth-server/internal/auth/models"
	"go-auth-server/internal/auth/repo/user"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	UserRepository user.UserRepository
}

func (usecase *AuthUsecase) generateTokenLifetime() (issuedAt, expiresAt int64) {
	now := time.Now()

	issuedAt = now.Unix()
	expiresAt = now.Add(time.Hour * 72).Unix()
	return
}

func (usecase *AuthUsecase) createSession(user *models.User) *models.Session {
	issuedAt, expiresAt := usecase.generateTokenLifetime()

	return models.NewSession(user.Name, user.Email, issuedAt, expiresAt)
}

func (usecase *AuthUsecase) SignUp(name, email, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		return "", auth.ErrInvalidPasswordFormat
	}

	user, userErr := models.NewUser(name, email, string(hashedPassword))

	if userErr != nil {
		return "", userErr
	}

	repoErr := usecase.UserRepository.Add(user)

	if repoErr != nil {
		return "", repoErr
	}

	token, encodingErr := EncodeSession(usecase.createSession(user))

	if encodingErr != nil {
		return "", encodingErr
	}

	return token, nil
}

func (usecase *AuthUsecase) SignIn(email string, password string) (string, error) {
	user, fetchErr := usecase.UserRepository.GetByEmail(email)

	if fetchErr != nil {
		return "", fetchErr
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", auth.ErrInvalidPassword
	}

	token, encodingErr := EncodeSession(usecase.createSession(user))

	if encodingErr != nil {
		return "", encodingErr
	}

	return token, nil
}

func (usecase *AuthUsecase) IsSessionExpired(session *models.Session) bool {
	return session.ExpiresAt < time.Now().Unix()
}

func (usecase *AuthUsecase) Rewoke(token string) (*models.Session, error) {
	session, decodingErr := DecodeSession(token)

	if decodingErr != nil {
		return nil, decodingErr
	}

	if usecase.IsSessionExpired(session) {
		return nil, auth.ErrExpiredToken
	}

	updatedUser, fetchErr := usecase.UserRepository.GetByEmail(session.Email)

	if fetchErr != nil {
		return nil, fetchErr
	}

	return models.NewSession(
		updatedUser.Name,
		updatedUser.Email,
		session.IssuedAt,
		session.ExpiresAt,
	), nil
}

func (usecase *AuthUsecase) RenewSession(session *models.Session) (string, error) {
	return EncodeSession(session)
}

func (usecase *AuthUsecase) RenewSessionWithToken(authorization string) (string, error) {
	session, decodingErr := DecodeSession(authorization)

	if decodingErr != nil {
		return "", decodingErr
	}

	if !usecase.IsSessionExpired(session) {
		return authorization, nil
	}

	return usecase.RenewSession(session)
}
