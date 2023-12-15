package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/common"
	"gorm.io/gorm"
)

type User struct {
    common.Model
	ExternalId  string `gorm:"not null;unique"json:"external_id"`
	MoneyAmount uint   `gorm:"not null;default:0"json:"money_amount"`
}

func Index(c *gin.Context, db *gorm.DB) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Return the user list
	c.JSON(200, users)
}
