package database

import (
	"time"

	"github.com/frtasoniero/subsmanager/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSampleProducts returns sample product data
func GetSampleProducts() []interface{} {
	return []interface{}{
		entities.Product{
			Name:        "Netflix",
			Description: "Streaming service with movies and TV shows",
			Price:       15.99,
			BillingType: "monthly",
			Category:    "streaming",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		entities.Product{
			Name:        "Spotify",
			Description: "Music streaming platform",
			Price:       9.99,
			BillingType: "monthly",
			Category:    "music",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		entities.Product{
			Name:        "Disney+",
			Description: "Disney streaming platform",
			Price:       7.99,
			BillingType: "monthly",
			Category:    "streaming",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		entities.Product{
			Name:        "Amazon Prime",
			Description: "Amazon Prime membership",
			Price:       12.99,
			BillingType: "monthly",
			Category:    "shopping",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		entities.Product{
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
		entities.User{
			Username:  "admin",
			Password:  "admin123", // In production, this should be hashed
			Email:     "admin@example.com",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		entities.User{
			Username:  "john_doe",
			Password:  "password123", // In production, this should be hashed
			Email:     "john@example.com",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		entities.User{
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
		entities.Subscription{
			UserID:       userIDs[1],    // john_doe
			ProductID:    productIDs[0], // Netflix
			Status:       "active",
			StartDate:    now.AddDate(0, -2, 0), // Started 2 months ago
			NextBilling:  now.AddDate(0, 1, 0),  // Next billing in 1 month
			PriceAtStart: 15.99,
			CreatedAt:    now.AddDate(0, -2, 0),
			UpdatedAt:    now,
		},
		entities.Subscription{
			UserID:       userIDs[1],    // john_doe
			ProductID:    productIDs[1], // Spotify
			Status:       "active",
			StartDate:    now.AddDate(0, -1, 0), // Started 1 month ago
			NextBilling:  now.AddDate(0, 1, 0),  // Next billing in 1 month
			PriceAtStart: 9.99,
			CreatedAt:    now.AddDate(0, -1, 0),
			UpdatedAt:    now,
		},
		entities.Subscription{
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
