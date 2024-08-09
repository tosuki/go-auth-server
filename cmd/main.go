package main

import (
	"auth-server/cmd/http"
	"auth-server/internal/auth/repo/users"
	"auth-server/internal/auth/usecase"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env") //default load .env file
	repositoryAdapter, err := ConnectDatabase()

	if err != nil {
		fmt.Printf("Failed to connect to the database, %s", err)
		return
	}

	authUsecase := usecase.AuthUsecase{
		Repository: users.NewUserRepository(repositoryAdapter),
	}

	http.RunServer(&authUsecase)
}
