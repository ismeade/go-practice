package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/", "D:/Games/2706952950")
	// r.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "HelloWorld!",
	// 	})
	// })
	r.Run()
}
