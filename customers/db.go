package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// This file is kept for compatibility reasons.
// All database-related functions have been moved to migrations.go.

// InitDB initializes the database connection
func InitDB() error {
	connStr := os.Getenv("DB_CONNECTION_STRING")

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database connection: %v", err)
	}

	fmt.Println("DB_CONNECTION_STRING:", connStr)

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return nil
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
