package admins

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// Private
	adminRouterPrivate := Router.Group("/admin")
	// // userRouterPrivate.Use(Limiter())
	// // userRouterPrivate.Use(Authen())
	// // userRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/activate")
	}
}
