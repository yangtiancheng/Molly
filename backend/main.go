package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Go!",
		})
	})
	r.Run(":8081") // 默认监听 0.0.0.0:8081
}
