package usecase

import (
	"errors"
	"go-auth-server/internal/auth"
	"go-auth-server/internal/auth/models"
	"log/slog"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var token_secret_key []byte = []byte(os.Getenv("JWT_SECRET"))

func EncodeSession(session *models.Session) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":      session.Name,
		"email":     session.Email,
		"issuedAt":  session.IssuedAt,
		"expiresAt": session.ExpiresAt,
	})

	signedToken, err := token.SignedString(token_secret_key)

	if err != nil {
		slog.Error("Failed to encode the session", "err", err.Error())
		return "", auth.ErrFailedToEncodeSession
	}

	return signedToken, nil
}

func DecodeSession(authorizationToken string) (*models.Session, error) {
	decoded, err := jwt.ParseWithClaims(authorizationToken, &models.Session{}, func(t *jwt.Token) (interface{}, error) {
		return token_secret_key, nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, auth.ErrInvalidToken
		}
		return nil, err
	}

	claims := decoded.Claims.(*models.Session)

	return claims, nil
}
