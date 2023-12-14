package betsybackend

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        fmt.Printf("%+v\n", c)
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    router.Run()
}
