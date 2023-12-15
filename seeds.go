package main

import (
	"fmt"
	"math/rand"

	"github.com/ventcode/betsy-backend/models"
	"github.com/ventcode/betsy-backend/user"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []user.User{
		{ExternalId: "greatGoogleId1", MoneyAmount: 1000},
		{ExternalId: "greatGoogleId2", MoneyAmount: 1000},
		{ExternalId: "greatGoogleId3", MoneyAmount: 1000},
		{ExternalId: "greatGoogleId4", MoneyAmount: 1000},
		{ExternalId: "greatGoogleId5", MoneyAmount: 1000},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			fmt.Printf("Error inserting user with ExternalId %s: %v\n", user.ExternalId, result.Error)
			continue
		}
	}
}

func SeedChallenges(db *gorm.DB) {
	var users []user.User
	db.Order("RANDOM()").Limit(5).Find(&users)

	for i, challenger := range users {
		var challenged user.User
		if i == len(users)-1 {
			challenged = users[0]
		} else {
			challenged = users[i+1]
		}

		challenge := models.Challenge{Challenger: challenger, Challenged: challenged, Title: "Great challenge"}
		result := db.Create(&challenge)
		if result.Error != nil {
			fmt.Printf("Error inserting challenge with Title %s: %v\n", challenge.Title, result.Error)
			continue
		}
	}
}

func SeedBets(db *gorm.DB) {
	var users []user.User
	db.Order("RANDOM()").Limit(5).Find(&users)

	var challenges []models.Challenge
	db.Order("RANDOM()").Limit(5).Find(&challenges)

	for _, user := range users {
		challenge := challenges[rand.Intn(len(challenges)-1)]
		bet := models.Bet{User: user, Challenge: challenge, BetOnChallenger: false, Amount: uint(rand.Uint32())}
		result := db.Create(&bet)
		if result.Error != nil {
			fmt.Println(result.Error)
			continue
		}
	}
}
