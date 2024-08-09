package auth

import "fmt"

var (
	ErrorInvalidEmail       = fmt.Errorf("invalid-email")
	ErrorInvalidPassword    = fmt.Errorf("invalid-password")
	ErrorTokenCreationError = fmt.Errorf("token-creation-err")
)
