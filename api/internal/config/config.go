package config

import (
	"log"
	"os"
	"time"
)

// Config holds all configuration for the application
type Config struct {
	// Database configuration
	Database DatabaseConfig `json:"database"`

	// Server configuration
	Server ServerConfig `json:"server"`
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	URI     string        `json:"uri"`
	Name    string        `json:"name"`
	Timeout time.Duration `json:"timeout"`
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string `json:"port"`
}

// Load loads configuration from environment variables with fallback defaults
func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			URI:     getEnv("MONGO_URI", "mongodb://root:password@localhost:27017"),
			Name:    getEnv("MONGO_DB_NAME", "subs-db"),
			Timeout: getDurationEnv("MONGO_TIMEOUT", 10*time.Second),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getDurationEnv gets a duration from environment variable with fallback default
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
		log.Printf("Warning: Invalid duration format for %s, using default", key)
	}
	return defaultValue
}
