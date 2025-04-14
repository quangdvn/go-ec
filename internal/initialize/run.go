package initialize

func Run() {
	LoadConfig()

	InitLogger()

	// InitMySql()

	InitMySqlC()

	InitServiceInterfaces()

	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
