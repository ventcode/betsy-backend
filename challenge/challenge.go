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
    Challenger user.User 
    Challenged user.User
    Title string `gorm:"not null"`
    Amount uint `gorm:"not null;default:0"`
    Status Status `gorm:"not null;default:0"`
    ChallengerWon *bool
}
