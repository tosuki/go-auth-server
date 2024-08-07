package usecase

import (
	"auth-server/internal/auth/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateLifetime(expirationTime time.Duration) (issuedAt, expiresAt int64) {
	now := time.Now()

	issuedAt = now.Unix()
	expiresAt = now.Add(time.Hour * expirationTime).Unix()

	return
}

func CreateToken(name, email, secretKey string) (string, error) {
	issuedAt, expiresAt := CreateLifetime(72)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  name,
		"email":     email,
		"issuedAt":  issuedAt,
		"expiresAt": expiresAt,
	})

	jwt, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return jwt, nil
}

func DecodeToken(tokenString, secretKey string) (*model.Session, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Session{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	var claims *model.Session = token.Claims.(*model.Session)

	return claims, nil
}

func IsValidToken(session *model.Session) bool {
	return session.ExpiresAt >= time.Now().Unix()
}
