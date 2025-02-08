package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("Server mode: ", global.Config.Server.Mode)
		r = gin.New()
		// r.Use(gin.Logger())
		// r.Use(gin.Recovery())
		// r.Use(middlewares.Cors())
		// r.Use(middlewares.RequestID())
		// r.Use(middlewares.Timeout())
		// r.Use(middlewares.RateLimit())
	}

	// Middlewares
	// r.Use() // Logger
	// r.Use() // CORS
	// r.Use() // Limit Global
	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2025")
	{
		MainGroup.GET("/healthCheck")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)

	}
	{
		adminRouter.InitUserRouter(MainGroup)
		adminRouter.InitAdminRouter(MainGroup)
	}
	return r
}
