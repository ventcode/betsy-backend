package challenge

import (
	"fmt"
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
	userid := c.Param("userid")
	message := "userid is " + userid
	c.String(http.StatusOK, message)
	fmt.Println(message)
}

func Update(c *gin.Context, db *gorm.DB) {
	cId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(422, "Cannot parse given ID")
		return
	}

	chall := Challenge{}
	ra := db.Preload("Challenger").Preload("Challenged").Find(&chall, cId).RowsAffected
	if ra == 0 {
		c.JSON(422, "Challenge not found")
		return
	}

	c.JSON(200, chall)
}

func Index(c *gin.Context, db *gorm.DB) {
	fmt.Println("Super")
}

type CreateChallengeInput struct {
	Title        string `json:"title"binding:"required"`
	Amount       uint   `json:"amount"binding:"required"`
	ChallengerID int    `json:"challenger_id"binding:"required"`
	ChallengedID int    `json:"challenged_id"binding:"required"`
}

func Create(c *gin.Context, db *gorm.DB) {
	var input CreateChallengeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", input)

	challenge := Challenge{Title: input.Title, Amount: input.Amount, ChallengerID: input.ChallengerID, ChallengedID: input.ChallengedID}
	db.Create(&challenge)

	fmt.Println(challenge)

	c.JSON(http.StatusOK, challenge)
}
