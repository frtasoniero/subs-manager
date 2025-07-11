package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user document
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
	Email     string             `bson:"email" json:"email"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// Product represents a product/service document
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	BillingType string             `bson:"billing_type" json:"billing_type"` // monthly, yearly, etc.
	Category    string             `bson:"category" json:"category"`         // streaming, music, etc.
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// Subscription represents a user's subscription to a product
type Subscription struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductID    primitive.ObjectID `bson:"product_id" json:"product_id"`
	Status       string             `bson:"status" json:"status"` // active, cancelled, expired
	StartDate    time.Time          `bson:"start_date" json:"start_date"`
	EndDate      *time.Time         `bson:"end_date,omitempty" json:"end_date,omitempty"`
	NextBilling  time.Time          `bson:"next_billing" json:"next_billing"`
	PriceAtStart float64            `bson:"price_at_start" json:"price_at_start"` // Price when subscription started
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

// UserSubscriptionView represents a joined view of user subscription with product details
type UserSubscriptionView struct {
	ID          primitive.ObjectID `json:"id"`
	UserID      primitive.ObjectID `json:"user_id"`
	ProductID   primitive.ObjectID `json:"product_id"`
	ProductName string             `json:"product_name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Status      string             `json:"status"`
	StartDate   time.Time          `json:"start_date"`
	EndDate     *time.Time         `json:"end_date,omitempty"`
	NextBilling time.Time          `json:"next_billing"`
	CreatedAt   time.Time          `json:"created_at"`
}

// GetSampleProducts returns sample product data
func GetSampleProducts() []interface{} {
	return []interface{}{
		Product{
			Name:        "Netflix",
			Description: "Streaming service with movies and TV shows",
			Price:       15.99,
			BillingType: "monthly",
			Category:    "streaming",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		Product{
			Name:        "Spotify",
			Description: "Music streaming platform",
			Price:       9.99,
			BillingType: "monthly",
			Category:    "music",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		Product{
			Name:        "Disney+",
			Description: "Disney streaming platform",
			Price:       7.99,
			BillingType: "monthly",
			Category:    "streaming",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		Product{
			Name:        "Amazon Prime",
			Description: "Amazon Prime membership",
			Price:       12.99,
			BillingType: "monthly",
			Category:    "shopping",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		Product{
			Name:        "YouTube Premium",
			Description: "Ad-free YouTube experience",
			Price:       11.99,
			BillingType: "monthly",
			Category:    "streaming",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
}

// GetSampleUsers returns sample user data
func GetSampleUsers() []interface{} {
	return []interface{}{
		User{
			Username:  "admin",
			Password:  "admin123", // In production, this should be hashed
			Email:     "admin@example.com",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		User{
			Username:  "john_doe",
			Password:  "password123", // In production, this should be hashed
			Email:     "john@example.com",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		User{
			Username:  "jane_smith",
			Password:  "password456", // In production, this should be hashed
			Email:     "jane@example.com",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// GetSampleSubscriptions returns sample subscription data
// Note: This would typically be called after users and products are created
func GetSampleSubscriptions(userIDs []primitive.ObjectID, productIDs []primitive.ObjectID) []interface{} {
	now := time.Now()
	return []interface{}{
		Subscription{
			UserID:       userIDs[1],    // john_doe
			ProductID:    productIDs[0], // Netflix
			Status:       "active",
			StartDate:    now.AddDate(0, -2, 0), // Started 2 months ago
			NextBilling:  now.AddDate(0, 1, 0),  // Next billing in 1 month
			PriceAtStart: 15.99,
			CreatedAt:    now.AddDate(0, -2, 0),
			UpdatedAt:    now,
		},
		Subscription{
			UserID:       userIDs[1],    // john_doe
			ProductID:    productIDs[1], // Spotify
			Status:       "active",
			StartDate:    now.AddDate(0, -1, 0), // Started 1 month ago
			NextBilling:  now.AddDate(0, 1, 0),  // Next billing in 1 month
			PriceAtStart: 9.99,
			CreatedAt:    now.AddDate(0, -1, 0),
			UpdatedAt:    now,
		},
		Subscription{
			UserID:       userIDs[2],    // jane_smith
			ProductID:    productIDs[0], // Netflix
			Status:       "active",
			StartDate:    now.AddDate(0, -3, 0), // Started 3 months ago
			NextBilling:  now.AddDate(0, 1, 0),  // Next billing in 1 month
			PriceAtStart: 15.99,
			CreatedAt:    now.AddDate(0, -3, 0),
			UpdatedAt:    now,
		},
	}
}
