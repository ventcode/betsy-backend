package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ExternalId  string `gorm:"not null;unique"`
	MoneyAmount uint   `gorm:"not null;default:0"`
}

func Index(c *gin.Context) {
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

	// Query the database
	var users []User
	result := gormDB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Return the user list
	c.JSON(200, users)
}
