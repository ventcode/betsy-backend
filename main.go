package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/bet"
	"github.com/ventcode/betsy-backend/challenge"
	"github.com/ventcode/betsy-backend/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    dsn := "user=postgres password=password dbname=postgres host=database port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    err = db.AutoMigrate(&user.User{}, &challenge.Challenge{}, &bet.Bet{})
    if err != nil {
        log.Fatal(err)
    }

    router := gin.Default()

    router.Run()
}
