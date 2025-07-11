package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	// Create database
	db := client.Database("subs-db")

	// Initialize collections in order (products, users, then subscriptions)
	productIDs, err := initializeProducts(ctx, db)
	if err != nil {
		return err
	}

	userIDs, err := initializeUsers(ctx, db)
	if err != nil {
		return err
	}

	if err := initializeSubscriptions(ctx, db, userIDs, productIDs); err != nil {
		return err
	}

	fmt.Println("✅ Database 'subs-db' initialized with sample data!")
	return nil
}

func initializeProducts(ctx context.Context, db *mongo.Database) ([]primitive.ObjectID, error) {
	collection := db.Collection("products")

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("failed to count products: %w", err)
	}

	if count > 0 {
		fmt.Printf("⚠️  Products collection already contains %d documents. Skipping initialization.\n", count)
		return nil, nil
	}

	// Insert sample products
	sampleProducts := GetSampleProducts()
	result, err := collection.InsertMany(ctx, sampleProducts)
	if err != nil {
		return nil, fmt.Errorf("failed to insert products: %w", err)
	}

	// Extract inserted IDs
	var productIDs []primitive.ObjectID
	for _, id := range result.InsertedIDs {
		productIDs = append(productIDs, id.(primitive.ObjectID))
	}

	fmt.Printf("✅ Products collection created with %d sample records!\n", len(sampleProducts))
	return productIDs, nil
}

func initializeUsers(ctx context.Context, db *mongo.Database) ([]primitive.ObjectID, error) {
	collection := db.Collection("users")

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("failed to count users: %w", err)
	}

	if count > 0 {
		fmt.Printf("⚠️  Users collection already contains %d documents. Skipping initialization.\n", count)
		return nil, nil
	}

	// Insert sample users
	sampleUsers := GetSampleUsers()
	result, err := collection.InsertMany(ctx, sampleUsers)
	if err != nil {
		return nil, fmt.Errorf("failed to insert users: %w", err)
	}

	// Extract inserted IDs
	var userIDs []primitive.ObjectID
	for _, id := range result.InsertedIDs {
		userIDs = append(userIDs, id.(primitive.ObjectID))
	}

	fmt.Printf("✅ Users collection created with %d sample records!\n", len(sampleUsers))
	return userIDs, nil
}

func initializeSubscriptions(ctx context.Context, db *mongo.Database, userIDs, productIDs []primitive.ObjectID) error {
	collection := db.Collection("subscriptions")

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, map[string]any{})
	if err != nil {
		return fmt.Errorf("failed to count subscriptions: %w", err)
	}

	if count > 0 {
		fmt.Printf("⚠️  Subscriptions collection already contains %d documents. Skipping initialization.\n", count)
		return nil
	}

	// Only create subscriptions if we have users and products
	if len(userIDs) == 0 || len(productIDs) == 0 {
		fmt.Println("⚠️  No users or products found. Skipping subscription initialization.")
		return nil
	}

	// Insert sample subscriptions
	sampleSubscriptions := GetSampleSubscriptions(userIDs, productIDs)
	_, err = collection.InsertMany(ctx, sampleSubscriptions)
	if err != nil {
		return fmt.Errorf("failed to insert subscriptions: %w", err)
	}

	fmt.Printf("✅ Subscriptions collection created with %d sample records!\n", len(sampleSubscriptions))
	return nil
}
