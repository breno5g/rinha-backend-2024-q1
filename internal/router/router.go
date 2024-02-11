package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// apiPort := os.Getenv("API_PORT")
	apiPort := 8080
	r.Run(fmt.Sprintf(":%d", apiPort))
}
