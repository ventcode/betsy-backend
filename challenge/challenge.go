package challenge

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
	gorm.Model
	ChallengerID  int `gorm:"not null"`
	Challenger    user.User
	ChallengedId  int `gorm:"not null"`
	Challenged    user.User
	Title         string `gorm:"not null"`
	Amount        uint   `gorm:"not null;default:0"`
	Status        Status `gorm:"not null;default:0"`
	ChallengerWon *bool
}

func Show(c *gin.Context) {
	db, exists := c.Get("db")
	if !exists {
		c.JSON(500, gin.H{"error": "Failed to get database instance"})
		return
	}

	gormDB, ok := db.(*gorm.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid database instance type"})
		return
	}

	var cha Challenge

	if !ok {
		c.JSON(500, gin.H{"error": "Invalid database instance type"})
		return
	}

	id, _ := c.Params.Get("id")
	gormDB.Find(&cha, id)

	if (cha == Challenge{}) {
		c.JSON(404, gin.H{"error": "Not found"})
	} else {
		c.JSON(200, gin.H{"challenge": cha})
	}
}

func Index(c *gin.Context) {
	fmt.Println("Super")
}
