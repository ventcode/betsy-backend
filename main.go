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

// App Config
var appConfig = SetAppConfig()

func main() {
	// Database
	db := DatabaseConnection()
	err := db.AutoMigrate(&user.User{}, &challenge.Challenge{}, &bet.Bet{})
	if err != nil {
		log.Fatal(err)
	}

	// Seeds
	SeedUsers(db)
	SeedChallenges(db)
	SeedBets(db)

	// Router
	router := gin.Default()

	// Middleware
	router.Use(MiddlewareSetDB(db))

	// Routes
	router.GET("/users", useDB(user.Index))
	router.GET("/challenges", useDB(challenge.Index))
	router.GET("/challenges/:id", useDB(challenge.Show))
	// router.POST("/challenges", challenge.Create)
	router.PATCH("/challenges/:id", useDB(challenge.Update))
	// router.POST("/challenges/:id/bets", bet.Create)
	router.Run()
}

func useDB(controllerFunc func(*gin.Context, *gorm.DB)) func(*gin.Context) {
	return func(c *gin.Context) {
		gormDB := helpers.GetDB(c)

		controllerFunc(c, gormDB)
	}
}
