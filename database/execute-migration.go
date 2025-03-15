package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// ExecuteMigrationsUp execute migration up when start running the app
func ExecuteMigrationsUp() {
	m, err := migrate.New(
		"file://database/migrations",
		"postgres://postgres:postgres@localhost:5432/checkin?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
