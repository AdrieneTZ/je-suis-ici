package main

import (
	"flag"
	"je-suis-ici/config"
	"je-suis-ici/database"
	"log"
)

func main() {
	// Parse command-line arguments
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// Load configuration
	cfg := config.NewConfig()

	// Execute migrations if flag is set
	if *migrateFlag {
		log.Println("Running database migrations...")
		if err := database.ExecuteMigrationsUp(cfg); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migrations completed successfully")
		return
	}

	// Connect to database
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
}
