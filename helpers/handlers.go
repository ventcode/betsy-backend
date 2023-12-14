package helpers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDB(c *gin.Context) *gorm.DB {
	db, exists := c.Get("db")

	if !exists {
		c.JSON(500, gin.H{"error": "Failed to get database instance"})
		return nil
	}

	gormDB, ok := db.(*gorm.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid database instance type"})
		return nil
	}

	return gormDB
}
