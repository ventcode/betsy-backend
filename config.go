package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbHost     = "localhost"
	dbPort     = 5433
	dbUser     = "postgres"
	dbPassword = "password"
	dbName     = "postgres"
)

func DatabaseConnection() *gorm.DB {
	format := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	dsn := fmt.Sprintf(format, dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
