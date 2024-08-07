package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("dkaowpdkapwodk")

func CreateLifetime(expirationTime time.Duration) (issuedAt, expiresAt int64) {
	now := time.Now()

	issuedAt = now.Unix()
	expiresAt = now.Add(time.Hour * expirationTime).Unix()

	return
}

func CreateToken(name, email string) (string, error) {
	issuedAt, expiresAt := CreateLifetime(72)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  name,
		"email":     email,
		"issuedAt":  issuedAt,
		"expiresAt": expiresAt,
	})

	jwt, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return jwt, nil
}
