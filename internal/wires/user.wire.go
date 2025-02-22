//go:build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/quangdvn/go-ec/internal/controllers"
	"github.com/quangdvn/go-ec/internal/repositories"
	"github.com/quangdvn/go-ec/internal/services"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		controllers.NewUserController,
		services.NewUserService,
		repositories.NewUserRepository,
		repositories.NewAuthRepository,
	)

	return new(controllers.UserController), nil
}
