package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env") //default load .env file

	ConnectDatabase()
	RunServer()
}
