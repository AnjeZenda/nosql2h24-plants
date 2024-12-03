package main

import (
	"plants/internal/app"
	"plants/internal/config"
)

func main() {
	cfg := config.MustLoad()

	if err := app.Run(cfg); err != nil {
		panic(err)
	}
}

// {
//     "buyerId": "5f2d8c1d1d8e2f1a1a2b3c7c",
//     "plantId": "5f1d7c1d1d8e2f1a1a2b3c4e",
//     "price": 1500,
//     "sellerId": "5f2d8c1d1d8e2f1a1a2b3c7a",
//     "species": "Алоэ"
// }
