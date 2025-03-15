package database

import (
	"fmt"
	"je-suis-ici/config"
	"log"

	"github.com/jmoiron/sqlx"
)

// DB connection
type DB struct {
	*sqlx.DB
}

// New DB connection
func New(cfg *config.Config) (*DB, error) {
	// get connection string
	connString := cfg.DBConnString()
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("DB connection fail: %w", err)
	}

	// test connection with ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("DB ping fail: %w", err)
	}

	log.Println("DB connection success")
	return &DB{db}, nil
}

// Close DB connection
func (db *DB) Close() error {
	return db.DB.Close()
}
