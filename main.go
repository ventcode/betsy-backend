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

    router.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    router.GET("/users", user.Index)
    router.GET("/challenges", challenge.Index)
    router.GET("/challenges/:id", challenge.Show)
    router.POST("/challenges", challenge.Create)
    router.PATCH("/challenges", challenge.Update)
    router.POST("/challenges/:id/bets", bet.Create)

    router := gin.Default()

    router.Run()
}
