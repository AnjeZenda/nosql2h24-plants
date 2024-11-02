package main

import (
	"plants/internal/app"
	"plants/internal/config"
	"strconv"
)

func main() {
	cfg := config.MustLoad()
	app := app.New(strconv.Itoa(cfg.GRPC.Port))

	if err := app.Run(); err != nil {
		panic("AAA")
	}
}
