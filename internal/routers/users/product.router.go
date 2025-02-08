package users

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// Public
	productRouterPublic := Router.Group("/products")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
	// Private
}
