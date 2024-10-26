package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database
	err = InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer GetDB().Close()

	// Run migrations
	err = RunMigrations()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Start the server or run other application logic here
	// For example:
	// startServer()

	log.Println("Application started successfully")
}
