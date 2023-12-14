package challenge

import (
	"github.com/ventcode/betsy-backend/user"
	"gorm.io/gorm"
)

type Status int

const(
    New Status = iota
    Accepted
    Started
    Finished
    Rejected
)

type Challenge struct {
    gorm.Model
    ChallengerID int `gorm:"not null"`
    Challenger user.User 
    ChallengedId int `gorm:"not null"`
    Challenged user.User
    Title string `gorm:"not null"`
    Amount uint `gorm:"not null;default:0"`
    Status Status `gorm:"not null;default:0"`
    ChallengerWon *bool
}

func Show(c *gin.Context) {
    userid := c.Param("userid") 
    message := "userid is " + userid
    c.String(http.StatusOK, message)
    fmt.Println(message)
}
