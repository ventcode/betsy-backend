package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/bet"
	"github.com/ventcode/betsy-backend/challenge"
	"github.com/ventcode/betsy-backend/helpers"
	"github.com/ventcode/betsy-backend/user"
	"gorm.io/gorm"
)

func main() {
	// Database
	db := DatabaseConnection()
	err := db.AutoMigrate(&user.User{}, &challenge.Challenge{}, &bet.Bet{})
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(MiddlewareSetDB(db))

	router.GET("/users", useDB(user.Index))
	router.GET("/challenges", useDB(challenge.Index))
	router.GET("/challenges/:id", useDB(challenge.Show))
	router.POST("/challenges", useDB(challenge.Create))
	router.PATCH("/challenges/:id", useDB(challenge.Update))
	// router.POST("/challenges/:id/bets", bet.Create)

	// generateMockData(db)
	router.Run()
}

func useDB(controllerFunc func(*gin.Context, *gorm.DB)) func(*gin.Context) {
	return func(c *gin.Context) {
		gormDB := helpers.GetDB(c)

		controllerFunc(c, gormDB)
	}
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

	uu := &user.User{ExternalId: "google", MoneyAmount: 2000}
	handleError(db.Create(uu))

	ch := &challenge.Challenge{Challenger: *u, Challenged: *uu, Title: "Great challenge"}
	handleError(db.Create(ch))
}
