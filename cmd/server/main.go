package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping/:name", Pong)
		v1.GET("/pong", Pong)
	}
	r.Run(":8002")
}

func Pong(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "pong" + name,
	})
}
