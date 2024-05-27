package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(RequestInfo())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HelloWorld!",
		})
	})
	r.Run()
}

func RequestInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Copy().FullPath()
		method := ctx.Copy().Request.Method
		fmt.Println("请求path:", path)
		fmt.Println("请求method:", method)
	}

}
