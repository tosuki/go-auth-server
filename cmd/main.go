package main

import (
	"github.com/joho/godotenv"
    "auth-server/cmd/auth"
)

func main() {
	godotenv.Load() //default load .env file

	auth.RunServer()
}
