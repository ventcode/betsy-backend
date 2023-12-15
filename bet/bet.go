package bet

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/common"
	"github.com/ventcode/betsy-backend/models"
	"github.com/ventcode/betsy-backend/user"
	"gorm.io/gorm"
)

type Bet struct {
	common.Model
	UserID          int              `gorm:"not null"json:"-"`
	User            user.User        `json:"user"`
	ChallengeID     int              `gorm:"not null"json:"challenge_id"`
	Challenge       models.Challenge `json:"-"`
	BetOnChallenger bool             `gorm:"not null"json:"bet_on_challenger"`
	Amount          uint             `gorm:"not null;check:amount > 0"json:"amount"`
}

type BetCreate struct {
	common.Model
	UserID          int  `json:"user_id" binding:"required"`
	ChallengeID     int  `json:"challenge_id" binding:"required"`
	BetOnChallenger bool `json:"bet_on_challenger" binding:"required"`
	Amount          uint `json:"amount" binding:"required,gt=0"`
}

type Status int

const (
	New Status = iota
	Accepted
	Started
	Finished
	Rejected
)

type Challenge struct {
	ID uint `gorm:"primarykey" json:"id"`
	common.Model
	ChallengerID  int       `gorm:"not null"json:"-"`
	Challenger    user.User `json:"challenger"`
	ChallengedID  int       `gorm:"not null"json:"-"`
	Challenged    user.User `json:"challenged"`
	Title         string    `gorm:"not null"json:"title"`
	Amount        uint      `gorm:"not null;default:0"json:"amount"`
	Status        Status    `gorm:"not null;default:0"json:"status"`
	ChallengerWon *bool     `json:"challenger_won"`
}

type CreateRequest struct {
	ChallengeID       int   `json: challenge_id"`
	UserID            int   `json:"user_id"`
	BetOnChallengerID *bool `json:"bet_on_challenger_id"`
	Amount            int   `json:"amount"`
}

func Create(c *gin.Context, db *gorm.DB) {
	var bet BetCreate

	if err := c.ShouldBindJSON(&bet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing_bets []BetCreate

	db.Table("bets").Where(&Bet{ChallengeID: bet.ChallengeID, UserID: bet.UserID}).Find(&existing_bets)

	if len(existing_bets) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bet for this challenge already exists"})
		return
	}

	var chal models.Challenge

	rows_affected := db.Table("challenges").Where(Challenge{Status: 1, ID: uint(bet.ChallengeID)}).Find(&chal).RowsAffected

	err := db.Table("bets").Create(&bet).Error

	if rows_affected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find challenge that allows for betting"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"challenge": bet})
	}
}
