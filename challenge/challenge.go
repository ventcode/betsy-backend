package challenge

import (
	"encoding/json"
	"fmt"
	"io"
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

func Create(c *gin.Context, db *gorm.DB) {

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read request body"})
		return
	}

	// Create an instance of YourStruct
	var newChallenge Challenge

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(jsonData, &newChallenge); err != nil {
		c.JSON(400, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	var user user.User
	getFirstUserResult := db.First(&user)
	if getFirstUserResult.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch the first user"})
		return
	}

	newChallenge.ChallengedID = int(user.ID)
	newChallenge.ChallengerID = int(user.ID)
	// TODO: Use the status enum somehow
	newChallenge.Status = 0

	// Create the user in the database
	createChallengeResult := db.Create(&newChallenge)
	if createChallengeResult.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create challenge"})
		return
	}

	// Return the created user
	c.JSON(201, newChallenge)

}
