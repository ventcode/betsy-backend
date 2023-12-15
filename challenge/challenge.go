package challenge

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	userid := c.Param("userid")
	message := "userid is " + userid
	c.String(http.StatusOK, message)
	fmt.Println(message)
}

func Index(c *gin.Context) {
	fmt.Println("Super")
}

func Create(c *gin.Context) {
	// Access GORM DB instance from Gin's context
	db, exists := c.Get("db")
	if !exists {
		c.JSON(500, gin.H{"error": "Failed to get database instance"})
		return
	}

	// Type assertion to convert interface{} to *gorm.DB
	gormDB, ok := db.(*gorm.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid database instance type"})
		return
	}

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
	getFirstUserResult := gormDB.First(&user)
	if getFirstUserResult.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch the first user"})
		return
	}

	newChallenge.ChallengedId = int(user.ID)
	newChallenge.ChallengerID = int(user.ID)
	// TODO: Use the status enum somehow
	newChallenge.Status = 0

	// Create the user in the database
	createChallengeResult := gormDB.Create(&newChallenge)
	if createChallengeResult.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create challenge"})
		return
	}

	// Return the created user
	c.JSON(201, newChallenge)

}
