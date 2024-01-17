package user

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ventcode/betsy-backend/models"
	"gorm.io/gorm"
)

const googleUrl = "https://www.googleapis.com/oauth2/v3/userinfo"

type UserResponse struct {
	Sub string
}

func Index(c *gin.Context, db *gorm.DB) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Return the user list
	c.JSON(200, users)
}

func Show(c *gin.Context, db *gorm.DB) {
	token := c.Request.Header["Authorization"][0]
	request, _ := http.NewRequest("GET", googleUrl, nil)
	request.Header.Set("Authorization", token)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}

	defer response.Body.Close()
	var parsedResponse UserResponse
	if response.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&parsedResponse)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
			return
		}
	}

	user := models.User{}
	result := db.Find(&user, "external_id = ?", parsedResponse.Sub)

	if result.RowsAffected == 0 {
		user.ExternalId = parsedResponse.Sub
		user.MoneyAmount = 1000
		res := db.Create(&user)
		if res.Error != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Record Invalid"})
		}

		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetBets(c *gin.Context, db *gorm.DB) {
	var bets []models.Bet
	id, _ := c.Params.Get("id")

	db.Preload("Challenge").Preload("User").Where("user_id = ?", id).Find(&bets)

	c.JSON(http.StatusOK, gin.H{"data": bets})
}
