package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdvn/go-ec/internal/services"
	"github.com/quangdvn/go-ec/pkg/responses"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// Controller -> Service -> Repo -> Models -> Repo
func (uc *UserController) GetUserById(c *gin.Context) {
	// if err != nil {
	// 	responses.ErrorResponse(c, 20004)
	// }
	responses.SuccessResponse(c, 20001, uc.userService.GetUserService())
}
