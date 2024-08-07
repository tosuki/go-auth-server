package main

import (
	"auth-server/internal/auth/usecase"
	"fmt"
)

func main() {
	token, err := usecase.CreateToken("kdapowdk", "dkawdk@gmail.com")

	if err != nil {
		fmt.Println("Failed to create the token, here it is the error: ", err)
	}

	fmt.Println("token: ", token)
}