package auth

import "errors"

var (
	ErrorInvalidEmail       = errors.New("invalid-email")
	ErrorEmailTaken         = errors.New("email-taken")
	ErrorInvalidPassword    = errors.New("invalid-password")
	ErrorTokenCreationError = errors.New("token-creation-err")
	ErrorUnsupported        = errors.New("unsupported")
	ErrorUUIDCreation       = errors.New("uuid creation err")
)
