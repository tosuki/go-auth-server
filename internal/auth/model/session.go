package model

import "github.com/golang-jwt/jwt/v5"

type Session struct {
	Name  string
	Email string

	IssuedAt  int64
	ExpiresAt int64

	jwt.RegisteredClaims
}

func NewSession(name, email string, issuedAt, expiresAt int64) (*Session, error) {
	session := Session{
		Name:      name,
		Email:     email,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}

	return &session, nil
}
