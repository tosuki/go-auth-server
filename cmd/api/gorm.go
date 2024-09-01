package main

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase(connection string) *gorm.DB {
	os.Getenv("POSTGRES_DSN")

	adapter, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connection,
	}))

	if err != nil {
		panic(err)
	}

	return adapter
}
