package initialize

func Run() {
	LoadConfig()

	InitLogger()

	// InitMySql()

	InitMySqlC()

	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
