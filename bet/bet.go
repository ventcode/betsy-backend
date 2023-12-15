package bet

import (
	"github.com/ventcode/betsy-backend/challenge"
	"github.com/ventcode/betsy-backend/user"
	"github.com/ventcode/betsy-backend/common"
)

type Bet struct {
    common.Model
    UserID          int `gorm:"not null"json:"-"`
    User            user.User `json:"user"`
    ChallengeID     int `gorm:"not null"json:"challenge_id"`
    Challenge       challenge.Challenge `json:"-"`
    BetOnChallenger bool `gorm:"not null"json:"bet_on_challenger"`
    Amount          uint `gorm:"not null;check:amount > 0"json:"amount"`
}
