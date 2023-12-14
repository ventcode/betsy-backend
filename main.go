package main

import (
	"fmt"
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

    generateMockData(db)

    router := gin.Default()

    router.Run()
}

func generateMockData(db *gorm.DB) {
    handleError := func(tx *gorm.DB) {
        err := tx.Error
        if err != nil {
            log.Fatal(err)
        }
    }

    u := &user.User{ExternalId: "greatGoogleId", MoneyAmount: 1000}
    handleError(db.Create(u))
    
    uu := &user.User{ExternalId: "greatGoogleId", MoneyAmount: 2000}
    handleError(db.Create(uu))

    ch := &challenge.Challenge{Challenger: *u, Challenged: *uu, Title: "Great challenge"}
    handleError(db.Create(ch))
}
