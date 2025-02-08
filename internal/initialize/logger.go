package initialize

import (
	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/pkg/loggers"
)

func InitLogger() {
	global.Logger = loggers.NewLogger(global.Config.Logger)
}
