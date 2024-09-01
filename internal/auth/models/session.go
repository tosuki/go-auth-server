package models

import "github.com/golang-jwt/jwt/v5"

type Session struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	IssuedAt  int64  `json:"issuedAt"`
	ExpiresAt int64  `json:"expiresAt"`

	jwt.RegisteredClaims
}

func NewSession(name, email string, issuedAt, expiresAt int64) *Session {
	return &Session{
		Name:      name,
		Email:     email,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}
}
