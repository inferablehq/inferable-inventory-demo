package main

import (
	"fmt"
	"log"
)

// RunMigrations executes all database migrations
func RunMigrations() error {
	db := GetDB()
	if db == nil {
		return fmt.Errorf("database connection not initialized")
	}

	log.Println("Running migrations...")

	migrations := []string{
		createCustomersTable,
		createCustomerOrdersTable,
		createOrderItemsTable,
	}

	for i, migration := range migrations {
		log.Printf("Running migration %d...", i+1)
		_, err := db.Exec(migration)
		if err != nil {
			return fmt.Errorf("error running migration %d: %v", i+1, err)
		}
	}

	log.Println("Migrations completed successfully")
	return nil
}

const createCustomersTable = `
CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
`

const createCustomerOrdersTable = `
CREATE TABLE IF NOT EXISTS customer_orders (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);
`

const createOrderItemsTable = `
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES customer_orders(id)
);
`
