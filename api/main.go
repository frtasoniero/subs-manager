package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Subscription struct {
	ID          interface{} `json:"id" bson:"_id"`
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	Price       float64     `json:"price" bson:"price"`
	Billing     string      `json:"billing" bson:"billing"`
	Created     time.Time   `json:"created" bson:"created"`
}

var mongoClient *mongo.Client
var database *mongo.Database

func initMongoDB() {
	uri := "mongodb://root:password@localhost:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	mongoClient = client
	database = client.Database("subs-db")
	log.Println("✅ Connected to MongoDB!")
}

func getSubscriptions(c *gin.Context) {
	collection := database.Collection("subscriptions")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscriptions"})
		return
	}
	defer cursor.Close(ctx)

	var subscriptions []Subscription
	if err = cursor.All(ctx, &subscriptions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode subscriptions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  subscriptions,
		"count": len(subscriptions),
	})
}

func getSubscriptionByName(c *gin.Context) {
	name := c.Param("name")
	collection := database.Collection("subscriptions")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var subscription Subscription
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&subscription)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscription})
}

func main() {
	// Initialize MongoDB connection
	initMongoDB()
	defer mongoClient.Disconnect(context.TODO())

	r := gin.Default()

	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Subscription endpoints
	r.GET("/subscriptions", getSubscriptions)
	r.GET("/subscriptions/:name", getSubscriptionByName)

	log.Println("🚀 Server starting on :8080")
	r.Run() // listen and serve on 0.0.0.0:8080
}
