package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/bet"
	"github.com/ventcode/betsy-backend/challenge"
	"github.com/ventcode/betsy-backend/helpers"
	"github.com/ventcode/betsy-backend/models"
	"github.com/ventcode/betsy-backend/user"
	"gorm.io/gorm"
)

// App Config
var appConfig = SetAppConfig()

func main() {
	// Database
	db := DatabaseConnection()
	err := db.AutoMigrate(&models.User{}, &models.Challenge{}, &models.Bet{})
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
	router.GET("/user", useDB(user.Show))
	router.GET("/users", useDB(user.Index))
	router.GET("/users/:id/bets", useDB(user.GetBets))
	router.GET("/challenges", useDB(challenge.Index))
	router.GET("/challenges/:id", useDB(challenge.Show))
	router.POST("/challenges", useDB(challenge.Create))
	router.PATCH("/challenges/:id", useDB(challenge.Update))
	router.GET("/bets/:id", useDB(bet.Show))
	router.POST("/bets", useDB(bet.Create))
	router.Run()
}

func useDB(controllerFunc func(*gin.Context, *gorm.DB)) func(*gin.Context) {
	return func(c *gin.Context) {
		gormDB := helpers.GetDB(c)

		controllerFunc(c, gormDB)
	}
}
