package main

import (
	"fmt"
	"log"
	"os"

	"github.com/frtasoniero/subsmanager/database"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cli/main.go <command>")
		fmt.Println("Commands:")
		fmt.Println("  init-db    Initialize and populate database")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init-db":
		fmt.Println("🚀 Initializing MongoDB database...")
		if err := database.InitializeDatabase(); err != nil {
			log.Fatal("Failed to initialize database:", err)
		}
		fmt.Println("✅ Database initialized successfully!")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
