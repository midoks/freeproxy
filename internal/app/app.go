package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start(port int) {

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", port))
}
