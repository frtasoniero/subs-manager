package app

import (
	"log"
	"net/http"

	"github.com/frtasoniero/subsmanager/internal/infrastructure/web"
	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	app    *App
	router *gin.Engine
	port   string
}

// NewServer creates a new HTTP server instance
func NewServer(app *App) *Server {
	addr := ":" + app.Config.Server.Port

	server := &Server{
		app:  app,
		port: addr,
	}

	// Initialize router
	server.router = server.setupRouter()

	return server
}

// setupRouter configures the Gin router with middleware and routes
func (s *Server) setupRouter() *gin.Engine {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Create router
	r := gin.New()

	// Add middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Setup routes
	web.SetupRoutes(r, s.app.Handlers)

	return r
}

// Start starts the HTTP server
func (s *Server) Start() error {
	log.Printf("🚀 Server starting on %s", s.port)
	return s.router.Run(s.port)
}

// Handler returns the HTTP handler (useful for testing)
func (s *Server) Handler() http.Handler {
	return s.router
}
