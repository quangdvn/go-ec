package routers

import (
	"github.com/quangdvn/go-ec/internal/routers/admins"
	"github.com/quangdvn/go-ec/internal/routers/users"
)

type RouterGroup struct {
	User  users.UserRouterGroup
	Admin admins.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
