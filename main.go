package main

import (
	"flag"
	"je-suis-ici/api"
	"je-suis-ici/config"
	"je-suis-ici/database"
	"je-suis-ici/repositories"
	"je-suis-ici/services"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Parse command-line arguments
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	apiFlag := flag.Bool("api", false, "Run in API server mode")
	flag.Parse()

	// Load configuration
	cfg := config.NewConfig()

	// Connect to database
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Execute migrations if flag is set
	if *migrateFlag {
		log.Println("Running database migrations...")
		if err := database.ExecuteMigrationsUp(cfg); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migrations completed successfully")
		return
	}

	// Start API server if flag is set
	if *apiFlag {
		// Initialize repositories
		checkInRepo := repositories.NewCheckInRepository(db)

		// Initialize services
		checkInService := services.NewCheckInService(checkInRepo)

		// Initialize handlers
		handler := api.NewHandler(checkInService)

		// Set up routes
		router := api.SetupRoutes(handler)

		// Create server
		addr := getEnv("SERVER_ADDR", ":8080")
		server := &http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		}

		// Start server
		log.Printf("API server starting on %s", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// getEnv gets environment variables or returns default values
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
