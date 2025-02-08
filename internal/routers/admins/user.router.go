package admins

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public
	// userRouterPublic := Router.Group("/admin/user")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/otp")
	// }
	// Private
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/activate")
	}
}
