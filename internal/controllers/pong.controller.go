package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	fmt.Println("-------> Main controller")
	name := c.DefaultQuery("name", "quangdvn")
	uid := c.Query("uid")
	c.JSON(200, gin.H{
		"message": "pong " + name,
		"uid":     uid,
		"users":   []string{"user1", "user2", "user3"},
	})
}
