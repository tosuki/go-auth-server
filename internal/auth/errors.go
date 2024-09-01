package auth

import "errors"

var (
	ErrFailedToCreateUUID    = errors.New("failed-to-create-uuid")
	ErrFailedToEncodeSession = errors.New("failed-to-encode-session")
	ErrDuplicatedUser        = errors.New("duplicated-user")
	ErrInvalidPasswordFormat = errors.New("invalid-password-format")
	ErrInvalidEmail          = errors.New("invalid-email")
	ErrInvalidPassword       = errors.New("invalid-password")
	ErrExpiredToken          = errors.New("expired-token")
	ErrInvalidToken          = errors.New("invalid-token")
)
