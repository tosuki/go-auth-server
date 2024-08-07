package main

import (
	"auth-server/internal/auth/model"
	"auth-server/internal/auth/usecase"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() //default load .env file

	secretKey := os.Getenv("JWT_SECRET")

	if secretKey == "" {
		panic("Missing JWT_SECRET on env file")
	}

	user, userErr := model.NewUser("kdk", "aa", "kdoapwdk", "aa", 2323, 2323)

	if userErr != nil {
		fmt.Printf("failed to create the user")
		return
	}

	token, tokenErr := usecase.CreateToken(user.Name, user.Email, secretKey)

	if tokenErr != nil {
		fmt.Printf("failed to create the token, %s\n", tokenErr)
		return
	}

	fmt.Printf("The token: %s", token)

	tokenClaims, decodeErr := usecase.DecodeToken(token, secretKey)

	if decodeErr != nil {
		fmt.Printf("failed to decode the token, %s\n", decodeErr)
		return
	}

	fmt.Printf("Name: %s\nEmail: %s\n", tokenClaims.Name, tokenClaims.Email)
}
