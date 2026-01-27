package main

import (
	"log"
	"paybridge-transaction-service/internal/app"
	"paybridge-transaction-service/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	bootstrap := app.NewBootstrap(cfg)

	if err := bootstrap.Start(); err != nil {
		log.Fatalf("Startup error: %v", err)
	}
}
