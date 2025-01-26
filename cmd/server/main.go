package main

import (
	"github.com/quangdvn/go-ec/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8082")
}
