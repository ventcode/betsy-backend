package models

import (
	"github.com/ventcode/betsy-backend/common"
)

type ChallengeStatus int

const (
	New ChallengeStatus = iota
	Accepted
	Started
	Finished
	Rejected
)

type User struct {
	common.Model
	ExternalId  string `gorm:"not null;unique" json:"external_id"`
	MoneyAmount uint   `gorm:"not null;default:0" json:"money_amount"`
}

type Challenge struct {
	common.Model
	ChallengerID  int             `gorm:"not null" json:"-"`
	Challenger    User            `json:"challenger"`
	ChallengedID  int             `gorm:"not null" json:"-"`
	Challenged    User            `json:"challenged"`
	Title         string          `gorm:"not null" json:"title"`
	Amount        uint            `gorm:"not null;default:0" json:"amount"`
	Status        ChallengeStatus `gorm:"not null;default:0" json:"status"`
	ChallengerWon *bool           `json:"challenger_won"`
	Bets          []Bet           `json:"bets"`
}

type Bet struct {
	common.Model
	UserID          int       `gorm:"not null" json:"-"`
	User            User      `json:"user"`
	ChallengeID     int       `gorm:"not null" json:"challenge_id"`
	Challenge       Challenge `json:"-"`
	BetOnChallenger bool      `gorm:"not null" json:"bet_on_challenger"`
	Amount          uint      `gorm:"not null;check:amount > 0" json:"amount"`
}
