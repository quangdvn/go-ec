package main

import (
	"github.com/quangdvn/go-ec/internal/initialize"
	// "github.com/quangdvn/go-ec/internal/routers"
)

func main() {
	// r := routers.NewRouter()
	// initMySql()
	// initRedis()
	// initKafka()
	// r.Run(":8082")

	// Should only contain Run()
	initialize.Run()
}
