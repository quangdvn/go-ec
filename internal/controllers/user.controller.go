package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdvn/go-ec/internal/services"
	"github.com/quangdvn/go-ec/pkg/responses"
)

// type UserController struct {
// 	userService *services.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: services.NewUserService(),
// 	}
// }

// // Controller -> Service -> Repo -> Models -> Repo
// func (uc *UserController) GetUserById(c *gin.Context) {
// 	// if err != nil {
// 	// 	responses.ErrorResponse(c, 20004)
// 	// }
// 	responses.SuccessResponse(c, 20001, uc.userService.GetUserService())

// !! WITH INTERFACE
type UserController struct {
	userService services.IUserService
}

// NewUserController initializes UserController with IUserService
func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Controller -> Service -> Repo -> Models -> Repo
func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	responses.SuccessResponse(c, result, nil)
}
