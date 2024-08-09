package main

import (
	"auth-server/cmd/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env") //default load .env file

	ConnectDatabase()
	http.RunServer()
}
