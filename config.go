package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppConfig struct {
	DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DatabaseConfig.Host,
		appConfig.DatabaseConfig.Port,
		appConfig.DatabaseConfig.User,
		appConfig.DatabaseConfig.Password,
		appConfig.DatabaseConfig.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ENV(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func SetAppConfig() AppConfig {
	dbPort, _ := strconv.Atoi(ENV("DATABASE_PORT"))
	appConfig := AppConfig{
		DatabaseConfig: DatabaseConfig{
			Host:     ENV("DATABASE_HOST"),
			Port:     dbPort,
			User:     ENV("DATABASE_USER"),
			Password: ENV("DATABASE_PASSWORD"),
			Name:     ENV("DATABASE_NAME"),
		},
	}

	return appConfig
}
