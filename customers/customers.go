package main

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Customer structure
type Customer struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateCustomerParams structure
type CreateCustomerParams struct {
	Name  string
	Email string
}

// CreateCustomer function
func CreateCustomer(params CreateCustomerParams) (*Customer, error) {
	if params.Name == "" || params.Email == "" {
		return nil, errors.New("name and email are required")
	}

	customer := Customer{
		ID:        uuid.New().String(),
		Name:      params.Name,
		Email:     params.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := db.Exec("INSERT INTO customers (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		customer.ID, customer.Name, customer.Email, customer.CreatedAt, customer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// GetCustomerParams structure
type GetCustomerParams struct {
	ID string
}

// GetCustomer function
func GetCustomer(params GetCustomerParams) (*Customer, error) {
	var customer Customer
	err := db.QueryRow("SELECT id, name, email, created_at, updated_at FROM customers WHERE id = $1", params.ID).
		Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}
	return &customer, nil
}

// UpdateCustomerParams structure
type UpdateCustomerParams struct {
	ID    string
	Name  string
	Email string
}

// UpdateCustomer function
func UpdateCustomer(params UpdateCustomerParams) (*Customer, error) {
	customer, err := GetCustomer(GetCustomerParams{ID: params.ID})
	if err != nil {
		return nil, err
	}

	if params.Name != "" {
		customer.Name = params.Name
	}
	if params.Email != "" {
		customer.Email = params.Email
	}
	customer.UpdatedAt = time.Now()

	_, err = db.Exec("UPDATE customers SET name = $1, email = $2, updated_at = $3 WHERE id = $4",
		customer.Name, customer.Email, customer.UpdatedAt, customer.ID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// DeleteCustomerParams structure
type DeleteCustomerParams struct {
	ID string
}

// DeleteCustomer function
func DeleteCustomer(params DeleteCustomerParams) error {
	result, err := db.Exec("DELETE FROM customers WHERE id = $1", params.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("customer not found")
	}

	return nil
}

// ListCustomersParams structure
type ListCustomersParams struct {
	// You can add pagination or filtering parameters here if needed
}

// ListCustomers function
func ListCustomers(params ListCustomersParams) ([]Customer, error) {
	rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
