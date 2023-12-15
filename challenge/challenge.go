package challenge

import (
	"fmt"
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
		c.JSON(422, "Challenge not found")
		return
	} else {
		c.JSON(200, gin.H{"challenge": cha})
	}
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
