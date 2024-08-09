package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() //default load .env file

	ConnectDatabase()
	RunServer()
}
