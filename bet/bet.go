package bet

import (
	"github.com/ventcode/betsy-backend/challenge"
	"github.com/ventcode/betsy-backend/user"
	"gorm.io/gorm"
)

type Bet struct {
	gorm.Model
	UserID          int `gorm:"not null"`
	User            user.User
	ChallengeID     int `gorm:"not null"`
	Challenge       challenge.Challenge
	BetOnChallenger bool `gorm:"not null"`
	Amount          uint `gorm:"not null;check:amount > 0"`
}
