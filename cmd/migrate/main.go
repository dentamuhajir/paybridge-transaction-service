package main

import (
	"log"
	"os"
	"paybridge-transaction-service/internal/migrate"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go [up|down|status|redo]")
	}

	migrationType := os.Args[1]

	migrate.Run(migrationType)

}
