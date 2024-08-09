package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_DSN")

	if dsn == "" {
		panic("MISSING POSTGRES DSN")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("POSTGRES_DSN"),
	}))

	if err != nil {
		return nil, fmt.Errorf("couldn't connect to the database, %s", err)
	}

	return db, nil
}
