package main

import (
	"go-auth-server/cmd/api/http"
	"go-auth-server/internal/auth/repo/user"
	"go-auth-server/internal/auth/usecase"
)

func main() {
	InitializeEnv("POSTGRES_DSN", "JWT_SECRET")
	// adapter := InitializeDatabase(os.Getenv("POSTGRES_DSN"))

	// userRepository := user.NewUserRepository(adapter)
	authUsecase := &usecase.AuthUsecase{
		UserRepository: user.NewMockedUserRepository(),
	}

	http.InitializeServer(authUsecase)
}
