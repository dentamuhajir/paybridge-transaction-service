package main

import (
	"log"
	"paybridge-transaction-service/internal/config"
	"paybridge-transaction-service/internal/server"
)

// @title Paybridge Transaction Service API
// @version 1.0
// @description API documentation for Transaction Services
// @BasePath /api/v1
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	if err := server.Run(cfg); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
