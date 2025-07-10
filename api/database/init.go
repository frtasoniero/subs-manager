package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitializeDatabase creates the database and populates it with initial data
func InitializeDatabase() error {
	// MongoDB connection string
	uri := "mongodb://root:password@localhost:27017"

	// Create client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	defer client.Disconnect(context.TODO())

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	// Create database and collection
	db := client.Database("subs-db")
	collection := db.Collection("subscriptions")

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("failed to count documents: %w", err)
	}

	if count > 0 {
		fmt.Printf("⚠️  Database already contains %d documents. Skipping initialization.\n", count)
		return nil
	}

	// Sample subscriptions data
	sampleSubs := []interface{}{
		map[string]interface{}{
			"name":        "Netflix",
			"description": "Streaming service",
			"price":       15.99,
			"billing":     "monthly",
			"created":     time.Now(),
		},
		map[string]interface{}{
			"name":        "Spotify",
			"description": "Music streaming",
			"price":       9.99,
			"billing":     "monthly",
			"created":     time.Now(),
		},
	}

	_, err = collection.InsertMany(ctx, sampleSubs)
	if err != nil {
		return fmt.Errorf("failed to insert subscriptions: %w", err)
	}

	fmt.Println("✅ Database 'subs-db' initialized with sample data!")
	return nil
}
