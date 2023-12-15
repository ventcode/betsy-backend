package challenge

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/bet"
	"github.com/ventcode/betsy-backend/common"
	"github.com/ventcode/betsy-backend/user"
	"gorm.io/gorm"
)

type Status int

const (
	New Status = iota
	Accepted
	Started
	Finished
	Rejected
)

type Challenge struct {
    common.Model
    ChallengerID  int `gorm:"not null" json:"-"`
    Challenger    user.User `json:"challenger"`
	ChallengedID  int `gorm:"not null" json:"-"`
    Challenged    user.User `json:"challenged"`
    Title         string `gorm:"not null" json:"title"`
    Amount        uint   `gorm:"not null;default:0" json:"amount"`
    Status        Status `gorm:"not null;default:0" json:"status"`
    ChallengerWon *bool `json:"challenger_won"`
    Bets          []bet.Bet `json:"bets"`
}

type UpdateChallengeInput struct {
    Status        *Status `json:"status" binding:"required,max=3"`
    ChallengerWon *bool `json:"challenger_won"`
}

func Show(c *gin.Context, db *gorm.DB) {
	userid := c.Param("userid")
	message := "userid is " + userid
	c.String(http.StatusOK, message)
	fmt.Println(message)
}

func Update(c *gin.Context, db *gorm.DB) {
    chall := Challenge{}
    ra := db.Preload("Challenger").
        Preload("Challenged").
        Preload("Bets").
        Find(&chall, c.Param("id")).RowsAffected
    if ra == 0 {
        c.JSON(404, "Challenge not found")
        return
    }

    if chall.Status == Finished {
        c.JSON(422, "Challenge is Finished, can't change anything!")
        return
    }

    updateInput := UpdateChallengeInput{}
    err := c.ShouldBindJSON(&updateInput)
    if err != nil {
        c.JSON(422, err.Error())
        return
    }

    if *updateInput.Status != Finished  {
        if updateInput.ChallengerWon != nil {
            c.JSON(422, "You can't determine who won if it is not finished!")
            return
        }
    } else if updateInput.ChallengerWon == nil {
        c.JSON(422, "You need to specify challenger_won!")
        return
    }

    err = db.Model(&chall).Updates(updateInput).Error
    if err != nil {
        c.JSON(422, err)
        return
    }

    if chall.Status == Finished {
        db.Transaction(func(db *gorm.DB) error {
            // TODO: calculate sum of bets

            for i := range chall.Bets {
                if chall.Bets[i].BetOnChallenger && *chall.ChallengerWon {
                    // TODO: winner logic and check if .Bets[i].User is preloaded - I doubt it
                    // chall.Bets[i].User.MoneyAmount += chall.Bets[i].Amount
                }
                db.Updates(&chall.Bets[i].User)
            }

            return nil
        })
    }

    c.JSON(200, chall)
}

func Index(c *gin.Context, db *gorm.DB) {
	fmt.Println("Super")
}
