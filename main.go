package main

import (
	"log"

	"github.com/hossain-asif/hotel_auth_service/app"
	"github.com/hossain-asif/hotel_auth_service/config/env"
	"github.com/hossain-asif/hotel_auth_service/internal/router"
)

func main() {
	env.Load()

	cfg := app.NewConfig()
	application := app.NewApplication(cfg, router.Modules)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
