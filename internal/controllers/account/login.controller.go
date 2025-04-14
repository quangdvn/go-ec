package account

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdvn/go-ec/internal/services"
	"github.com/quangdvn/go-ec/pkg/responses"
)

var LoginController = new(cUserLogin)

type cUserLogin struct{}

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := services.UserLogin().Login(ctx)
	if err != nil {
		responses.ErrorResponse(ctx, responses.ErrCodeInvalidParam)
		return
	}
	responses.SuccessResponse(ctx, responses.ErrCodeSuccess, nil)
}
