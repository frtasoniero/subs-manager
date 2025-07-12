package app

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/frtasoniero/subsmanager/internal/config"
	"github.com/frtasoniero/subsmanager/internal/infrastructure/database"
	"github.com/frtasoniero/subsmanager/internal/infrastructure/database/repositories"
	"github.com/frtasoniero/subsmanager/internal/infrastructure/web"
	"github.com/frtasoniero/subsmanager/internal/infrastructure/web/handlers"
	"github.com/frtasoniero/subsmanager/internal/usecases"
)

// App represents the application container with all dependencies
type App struct {
	Config   *config.Config
	DB       *mongo.Database
	Client   *mongo.Client
	Handlers *web.AppHandlers
}

// NewApp creates a new application instance with all dependencies
func NewApp() (*App, error) {
	// Load configuration
	cfg := config.Load()
	log.Println("📋 Configuration loaded")

	// Database connection
	dbConfig := database.Config{
		URI:      cfg.Database.URI,
		Database: cfg.Database.Name,
		Timeout:  cfg.Database.Timeout,
	}

	client, db, err := database.NewConnection(dbConfig)
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	subscriptionRepo := repositories.NewMongoSubscriptionRepository(db)

	// Initialize use cases
	subscriptionUseCase := usecases.NewSubscriptionUseCase(subscriptionRepo)

	// Initialize handlers
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionUseCase)

	return &App{
		Config: cfg,
		DB:     db,
		Client: client,
		Handlers: &web.AppHandlers{
			Subscription: subscriptionHandler,
		},
	}, nil
}

// Close closes the application and cleans up resources
func (a *App) Close() error {
	return database.Close(a.Client)
}
