package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go [up|down|status|redo]")
	}

	migrationType := os.Args[1]

	cmd := exec.Command("sql-migrate", migrationType, "-config", "dbconfig.yml", "-env", "development")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Migration %s failed: %v", migrationType, err)
	}
	log.Printf("Migration %s completed successfully.", migrationType)
}
