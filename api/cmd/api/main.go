package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/frtasoniero/subsmanager/internal/infrastructure/database/repositories"
	"github.com/frtasoniero/subsmanager/internal/infrastructure/web/handlers"
	"github.com/frtasoniero/subsmanager/internal/usecases"
)

func main() {
	// Database connection
	uri := "mongodb://root:password@localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	db := client.Database("subs-db")
	log.Println("✅ Connected to MongoDB!")

	// Initialize repositories
	subscriptionRepo := repositories.NewMongoSubscriptionRepository(db)

	// Initialize use cases
	subscriptionUseCase := usecases.NewSubscriptionUseCase(subscriptionRepo)

	// Initialize handlers
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionUseCase)

	// Initialize router
	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		api.GET("/subscriptions", subscriptionHandler.GetAllSubscriptions)
	}

	// Start server
	log.Println("🚀 Server starting on :8080")
	r.Run(":8080")
}
