package challenge

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	ChallengerID  int       `gorm:"not null"json:"-"`
	Challenger    user.User `json:"challenger"`
	ChallengedID  int       `gorm:"not null"json:"-"`
	Challenged    user.User `json:"challenged"`
	Title         string    `gorm:"not null"json:"title"`
	Amount        uint      `gorm:"not null;default:0"json:"amount"`
	Status        Status    `gorm:"not null;default:0"json:"status"`
	ChallengerWon *bool     `json:"challenger_won"`
}

func Show(c *gin.Context, db *gorm.DB) {
	var cha Challenge

	id, _ := c.Params.Get("id")
	rows_affected := db.Find(&cha, id).RowsAffected

	if rows_affected == 0 {
		c.JSON(http.StatusUnprocessableEntity, "Challenge not found")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"challenge": cha})
	}
}

func Update(c *gin.Context, db *gorm.DB) {
	cId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Cannot parse given ID")
		return
	}

	chall := Challenge{}
	ra := db.Preload("Challenger").Preload("Challenged").Find(&chall, cId).RowsAffected
	if ra == 0 {
		c.JSON(http.StatusUnprocessableEntity, "Challenge not found")
		return
	}

	c.JSON(http.StatusOK, chall)
}

func Index(c *gin.Context, db *gorm.DB) {
	var challenges []Challenge
	db.Preload("Challenger").Preload("Challenged").Find(&challenges)

	c.JSON(http.StatusOK, gin.H{"data": challenges})
}

type CreateChallengeInput struct {
	Title        string `json:"title" binding:"required"`
	Amount       *uint  `json:"amount" binding:"required,gt=0"`
	ChallengerID *int   `json:"challenger_id" binding:"required"`
	ChallengedID *int   `json:"challenged_id" binding:"required"`
}

func NewChallenge(challInp *CreateChallengeInput, challenger, challenged *user.User) *Challenge {
	return &Challenge{
        Title: challInp.Title,
        Amount: *challInp.Amount,
        ChallengerID: *challInp.ChallengerID,
        Challenger: *challenger,
        ChallengedID: *challInp.ChallengedID,
        Challenged: *challenged,
    }
}

func Create(c *gin.Context, db *gorm.DB) {
	var input CreateChallengeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    challenger := &user.User{}
    ra := db.Find(challenger, input.ChallengerID).RowsAffected
    if ra == 0 {
        c.JSON(http.StatusNotFound, "Challenger user of provided specified id does not exist")
        return
    }

    challenged := &user.User{}
    ra = db.Find(challenged, input.ChallengedID).RowsAffected
    if ra == 0 {
        c.JSON(http.StatusNotFound, "Challenged user of provided specified id does not exist")
        return
    }

    if challenger.MoneyAmount < *input.Amount {
        c.JSON(http.StatusUnprocessableEntity, "Challenger has no enough money to start this challenge")
        return
    }

	challenge := NewChallenge(&input, challenger, challenged)
	db.Create(challenge)

	c.JSON(http.StatusOK, challenge)
}
