package users

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdvn/go-ec/internal/wires"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public
	// Non-Dependency
	// ur := repositories.NewUserRepository()
	// us := services.NewUserService(ur)
	// userHandlerNonDependency := controllers.NewUserController(us)

	// Wire - Dependency Injection
	// Detach high-level module from low-level module
	userController, _ := wires.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}
	// Private
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/profile")
	}
}
