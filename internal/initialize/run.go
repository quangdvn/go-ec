package initialize

import (
	"fmt"

	"github.com/quangdvn/go-ec/global"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading config MySQL", m.Username)

	InitLogger()
	global.Logger.Info("Logger initialized")

	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run(":8082")
}
