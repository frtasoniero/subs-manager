package web

import (
	"net/http"

	"github.com/frtasoniero/subsmanager/internal/infrastructure/web/handlers"
	"github.com/gin-gonic/gin"
)

// Handlers interface for dependency injection
type AppHandlers struct {
	Subscription *handlers.SubscriptionHandler
	// Add more handlers here as you create them
	// User         *handlers.UserHandler
	// Product      *handlers.ProductHandler
}

// SetupRoutes configures all routes for the application
func SetupRoutes(r *gin.Engine, appHandlers *AppHandlers) {
	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "healthy",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Subscription routes
		subscriptions := v1.Group("/subscriptions")
		{
			subscriptions.GET("", appHandlers.Subscription.GetAllSubscriptions)
			// Add other subscription routes as you implement them
			// subscriptions.GET("/:id", appHandlers.Subscription.GetSubscriptionByID)
			// subscriptions.POST("", appHandlers.Subscription.CreateSubscription)
			// subscriptions.PUT("/:id", appHandlers.Subscription.UpdateSubscription)
			// subscriptions.DELETE("/:id", appHandlers.Subscription.DeleteSubscription)
		}

		// User routes (placeholder for future implementation)
		users := v1.Group("/users")
		{
			users.GET("", func(c *gin.Context) {
				c.JSON(http.StatusNotImplemented, gin.H{"message": "Users endpoint not implemented yet"})
			})
			users.GET("/:id/subscriptions", func(c *gin.Context) {
				c.JSON(http.StatusNotImplemented, gin.H{"message": "User subscriptions endpoint not implemented yet"})
			})
		}

		// Product routes (placeholder for future implementation)
		products := v1.Group("/products")
		{
			products.GET("", func(c *gin.Context) {
				c.JSON(http.StatusNotImplemented, gin.H{"message": "Products endpoint not implemented yet"})
			})
		}
	}
}
