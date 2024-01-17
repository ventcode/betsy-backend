package bet

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/models"
	"gorm.io/gorm"
)

type BetCreate struct {
	UserID          int   `json:"user_id" binding:"required"`
	ChallengeID     int   `json:"challenge_id" binding:"required"`
	BetOnChallenger *bool `json:"bet_on_challenger" binding:"required"`
	Amount          uint  `json:"amount" binding:"required,gt=0"`
}

func Create(c *gin.Context, db *gorm.DB) {
	var bc BetCreate

	if err := c.ShouldBindJSON(&bc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := &models.Bet{}
	ra := db.Find(b, "challenge_id = ? AND user_id = ?", bc.ChallengeID, bc.UserID).RowsAffected
	if ra > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bet for this challenge already exists"})
		return
	}

	var chal models.Challenge
	rows_affected := db.Find(&chal, bc.ChallengeID).RowsAffected
	if rows_affected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find challenge that allows for betting"})
		return
	}

	usr := &models.User{}
	rows_affected = db.Find(&usr, bc.UserID).RowsAffected
	if rows_affected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find user"})
		return
	}

	bet := &models.Bet{
		UserID:          bc.UserID,
		User:            *usr,
		ChallengeID:     bc.ChallengeID,
		Amount:          bc.Amount,
		BetOnChallenger: *bc.BetOnChallenger,
	}
	err := db.Create(&bet).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, bet)
	}
}

func Show(c *gin.Context, db *gorm.DB) {
	var bet models.Bet

	id, _ := c.Params.Get("id")
	rows_affected := db.Preload("User").Preload("Challenge").Find(&bet, id).RowsAffected

	if rows_affected == 0 {
		c.JSON(http.StatusUnprocessableEntity, "Bet not found")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"bet": bet})
	}
}
