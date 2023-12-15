package bet

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/common"
	"github.com/ventcode/betsy-backend/models"
	"gorm.io/gorm"
)

type BetCreate struct {
	common.Model
	UserID          int  `json:"user_id" binding:"required"`
	ChallengeID     int  `json:"challenge_id" binding:"required"`
	BetOnChallenger bool `json:"bet_on_challenger" binding:"required"`
	Amount          uint `json:"amount" binding:"required,gt=0"`
}

func Create(c *gin.Context, db *gorm.DB) {
	var bet BetCreate

	if err := c.ShouldBindJSON(&bet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing_bets []models.Bet

	challenge_id, _ := c.Params.Get("challenge_id")
	user_id, _ := c.Params.Get("user_id")

	int_challenge_id, _ := strconv.Atoi(challenge_id)
	int_user_id, _ := strconv.Atoi(user_id)

	db.Where(&models.Bet{ChallengeID: int_challenge_id, UserID: int_user_id}).Find(&existing_bets)

	if len(existing_bets) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bet for this challenge already exists"})
		return
	}

	err := db.Table("bets").Create(&bet).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"challenge": bet})
	}
}
