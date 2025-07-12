package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config holds database connection configuration
type Config struct {
	URI      string
	Database string
	Timeout  time.Duration
}

// NewConnection creates a new MongoDB connection with the given configuration
func NewConnection(config Config) (*mongo.Client, *mongo.Database, error) {
	log.Printf("🔄 Connecting to MongoDB at %s...", config.URI)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.URI))
	if err != nil {
		log.Printf("❌ Failed to connect to MongoDB: %v", err)
		return nil, nil, err
	}

	// Test connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("❌ Failed to ping MongoDB: %v", err)
		return nil, nil, err
	}

	db := client.Database(config.Database)
	log.Printf("✅ Connected to MongoDB database: %s", config.Database)

	return client, db, nil
}

// Close closes the MongoDB connection
func Close(client *mongo.Client) error {
	log.Println("🔄 Closing MongoDB connection...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Printf("❌ Error closing MongoDB connection: %v", err)
		return err
	}

	log.Println("✅ MongoDB connection closed successfully")
	return nil
}
