package main

import (
	"log"

	"github.com/frtasoniero/subsmanager/internal/app"
)

func main() {
	// Initialize application with all dependencies
	application, err := app.NewApp()
	if err != nil {
		log.Fatal("Failed to initialize application:", err)
	}
	defer application.Close()

	// Create and configure server
	server := app.NewServer(application)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
