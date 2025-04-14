package initialize

import (
	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/internal/database"
	"github.com/quangdvn/go-ec/internal/services"
	"github.com/quangdvn/go-ec/internal/services/impls"
)

func InitServiceInterfaces() {
	queries := database.New(global.Mdbc)

	// User Service Interface
	services.InitUserLogin(impls.NewUserLoginImplement(queries))
}
