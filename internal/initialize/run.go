package initialize

func Run() {
	LoadConfig()

	InitLogger()

	InitMySql()

	InitRedis()

	r := InitRouter()

	r.Run(":8082")
}
