package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/quangdvn/go-ec/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserById)
	}
	return r
}
