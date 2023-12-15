package challenge

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/models"
	"gorm.io/gorm"
)

func Show(c *gin.Context, db *gorm.DB) {
	var cha models.Challenge

	id, _ := c.Params.Get("id")
	rows_affected := db.Find(&cha, id).RowsAffected

	if rows_affected == 0 {
		c.JSON(422, "Challenge not found")
		return
	} else {
		c.JSON(200, gin.H{"challenge": cha})
	}
}

// type UpdateChallengeInput struct {
// 	Status        *Status `json:"status" binding:"required,max=3"`
// 	ChallengerWon *bool   `json:"challenger_won"`
// 	common.Model
// 	ChallengerID  int       `gorm:"not null"json:"-"`
// 	Challenger    user.User `json:"challenger"`
// 	ChallengedID  int       `gorm:"not null"json:"-"`
// 	Challenged    user.User `json:"challenged"`
// 	Title         string    `gorm:"not null"json:"title"`
// 	Amount        uint      `gorm:"not null;default:0"json:"amount"`
// 	Status        Status    `gorm:"not null;default:0"json:"status"`
// 	ChallengerWon *bool     `json:"challenger_won"`
// }

// func Update(c *gin.Context, db *gorm.DB) {
// 	chall := Challenge{}
// 	ra := db.Preload("Challenger").
// 		Preload("Challenged").
// 		Preload("Bets").
// 		Find(&chall, c.Param("id")).RowsAffected
// 	if ra == 0 {
// 		c.JSON(404, "Challenge not found")
// 		return
// 	}

// 	if chall.Status == Finished {
// 		c.JSON(422, "Challenge is Finished, can't change anything!")
// 		return
// 	}

// 	updateInput := UpdateChallengeInput{}
// 	err := c.ShouldBindJSON(&updateInput)
// 	if err != nil {
// 		c.JSON(422, err.Error())
// 		return
// 	}

// 	if *updateInput.Status != Finished {
// 		if updateInput.ChallengerWon != nil {
// 			c.JSON(422, "You can't determine who won if it is not finished!")
// 			return
// 		}
// 	} else if updateInput.ChallengerWon == nil {
// 		c.JSON(422, "You need to specify challenger_won!")
// 		return
// 	}

// 	err = db.Model(&chall).Updates(updateInput).Error
// 	if err != nil {
// 		c.JSON(422, err)
// 		return
// 	}

// 	if chall.Status == Finished {
// 		db.Transaction(func(db *gorm.DB) error {
// 			// TODO: calculate sum of bets

// 			for i := range chall.Bets {
// 				if chall.Bets[i].BetOnChallenger && *chall.ChallengerWon {
// 					// TODO: winner logic and check if .Bets[i].User is preloaded - I doubt it
// 					// chall.Bets[i].User.MoneyAmount += chall.Bets[i].Amount
// 				}
// 				db.Updates(&chall.Bets[i].User)
// 			}

// 			return nil
// 		})
// 	}

// 	c.JSON(200, chall)
// 	cId, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(422, "Cannot parse given ID")
// 		return
// 	}

// 	chall := Challenge{}
// 	ra := db.Preload("Challenger").Preload("Challenged").Find(&chall, cId).RowsAffected
// 	if ra == 0 {
// 		c.JSON(422, "Challenge not found")
// 		return
// 	}

// 	c.JSON(200, chall)
// }

func Index(c *gin.Context, db *gorm.DB) {
	var challenges []models.Challenge
	db.Preload("Challenger").Preload("Challenged").Find(&challenges)

	c.JSON(http.StatusOK, gin.H{"data": challenges})
}

type CreateChallengeInput struct {
	Title        string `json:"title" binding:"required"`
	Amount       *uint  `json:"amount" binding:"required,gt=0"`
	ChallengerID *int   `json:"challenger_id" binding:"required"`
	ChallengedID *int   `json:"challenged_id" binding:"required"`
}

func Create(c *gin.Context, db *gorm.DB) {
	var input CreateChallengeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	challenge := models.Challenge{Title: input.Title, Amount: *input.Amount, ChallengerID: *input.ChallengerID, ChallengedID: *input.ChallengedID}
	db.Create(&challenge)

	c.JSON(http.StatusOK, challenge)
}

type UpdateChallengeInput struct {
	// ChallengeId models.Challenge `json:"challenge_id"`
	Status   models.Status `json:"status"`
	WinnerID uint          `json:"winner_id"`
}

func Update(c *gin.Context, db *gorm.DB) {

	//TODO: Update database on every status apart from Finished
	var input UpdateChallengeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var challenge models.Challenge
	id, _ := c.Params.Get("id")
	rows_affected := db.Find(&challenge, id).RowsAffected
	if rows_affected == 0 {
		c.JSON(422, "Challenge not found")
		return
	}

	if input.Status == models.Finished {

		challengerWon := challenge.ChallengerID == int(input.WinnerID)
		fmt.Println(challengerWon)

		//
		c.JSON(http.StatusOK, input)
		return
	}

	// challenge := models.Challenge{Title: input.Title, Amount: *input.Amount, ChallengerID: *input.ChallengerID, ChallengedID: *input.ChallengedID}
	// db.Create(&challenge)

	c.JSON(http.StatusOK, input)
}
