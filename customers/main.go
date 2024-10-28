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

	// Register Inferable functions
	err = registerInferableFunctions()
	if err != nil {
		log.Fatalf("Failed to register Inferable functions: %v", err)
	}

	log.Println("Application started successfully")

	// Keep the application running
	select {}
}
