package main

import (
	"fmt"
	"log"
	"os"

	"github.com/frtasoniero/subsmanager/internal/config"
	"github.com/frtasoniero/subsmanager/internal/infrastructure/database"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cli/main.go <command>")
		fmt.Println("Commands:")
		fmt.Println("  init-db    Initialize and populate database")
		fmt.Println("  clean-db   Clean database and reset to default data")
		os.Exit(1)
	}

	command := os.Args[1]

	// Load configuration
	cfg := config.Load()
	dbConfig := database.Config{
		URI:      cfg.Database.URI,
		Database: cfg.Database.Name,
		Timeout:  cfg.Database.Timeout,
	}

	switch command {
	case "init-db":
		fmt.Println("🚀 Initializing MongoDB database...")
		if err := database.InitializeDatabase(dbConfig); err != nil {
			log.Fatal("Failed to initialize database:", err)
		}
		fmt.Println("✅ Database initialized successfully!")

	case "clean-db":
		fmt.Println("🧹 Cleaning and resetting MongoDB database...")
		if err := database.CleanDatabase(dbConfig); err != nil {
			log.Fatal("Failed to clean database:", err)
		}
		fmt.Println("✅ Database cleaned and reset successfully!")

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
