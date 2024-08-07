package usecase

import (
    "auth-server/internal/auth/model"
)

type AuthenticationResponse struct {
    Err error
    Data *model.Session
}

func NewAuthenticationResponse() *AuthenticationResponse {
    return nil
}

func Authenticate(email, password string) (string, string) {
    return email, password
}
