package main

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// CustomerOrder structure
type CustomerOrder struct {
	ID         string
	CustomerID string
	Items      []struct {
		ProductID string
		Quantity  int
		Price     float64
	}
	TotalPrice float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// CreateCustomerOrderParams structure
type CreateCustomerOrderParams struct {
	CustomerID string
	Items      []struct {
		ProductID string
		Quantity  int
		Price     float64
	}
}

// CreateCustomerOrder function
func CreateCustomerOrder(params CreateCustomerOrderParams) (*CustomerOrder, error) {
	if params.CustomerID == "" || len(params.Items) == 0 {
		return nil, errors.New("customerID and items are required")
	}

	totalPrice := 0.0
	for _, item := range params.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	order := CustomerOrder{
		ID:         uuid.New().String(),
		CustomerID: params.CustomerID,
		Items:      params.Items,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO customer_orders (id, customer_id, total_price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		order.ID, order.CustomerID, order.TotalPrice, order.CreatedAt, order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	for _, item := range order.Items {
		_, err = tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
			order.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// GetCustomerOrderParams structure
type GetCustomerOrderParams struct {
	ID string
}

// GetCustomerOrder function
func GetCustomerOrder(params GetCustomerOrderParams) (*CustomerOrder, error) {
	var order CustomerOrder
	err := db.QueryRow("SELECT id, customer_id, total_price, created_at, updated_at FROM customer_orders WHERE id = $1", params.ID).
		Scan(&order.ID, &order.CustomerID, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer order not found")
		}
		return nil, err
	}

	rows, err := db.Query("SELECT product_id, quantity, price FROM order_items WHERE order_id = $1", params.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item struct {
			ProductID string
			Quantity  int
			Price     float64
		}
		err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}

	return &order, nil
}

// UpdateCustomerOrderParams structure
type UpdateCustomerOrderParams struct {
	ID    string
	Items []struct {
		ProductID string
		Quantity  int
		Price     float64
	}
}

// UpdateCustomerOrder function
func UpdateCustomerOrder(params UpdateCustomerOrderParams) (*CustomerOrder, error) {
	order, err := GetCustomerOrder(GetCustomerOrderParams{ID: params.ID})
	if err != nil {
		return nil, err
	}

	totalPrice := 0.0
	for _, item := range params.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE customer_orders SET total_price = $1, updated_at = $2 WHERE id = $3",
		totalPrice, time.Now(), params.ID)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = $1", params.ID)
	if err != nil {
		return nil, err
	}

	for _, item := range params.Items {
		_, err = tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
			params.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	order.Items = params.Items
	order.TotalPrice = totalPrice
	order.UpdatedAt = time.Now()

	return order, nil
}

// DeleteCustomerOrderParams structure
type DeleteCustomerOrderParams struct {
	ID string
}

// DeleteCustomerOrder function
func DeleteCustomerOrder(params DeleteCustomerOrderParams) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = $1", params.ID)
	if err != nil {
		return err
	}

	result, err := tx.Exec("DELETE FROM customer_orders WHERE id = $1", params.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("customer order not found")
	}

	return tx.Commit()
}

// ListCustomerOrdersParams structure
type ListCustomerOrdersParams struct {
	// You can add pagination or filtering parameters here if needed
}

// ListCustomerOrders function
func ListCustomerOrders(params ListCustomerOrdersParams) ([]CustomerOrder, error) {
	rows, err := db.Query("SELECT id, customer_id, total_price, created_at, updated_at FROM customer_orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []CustomerOrder
	for rows.Next() {
		var order CustomerOrder
		err := rows.Scan(&order.ID, &order.CustomerID, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// GetCustomerOrdersByCustomerIDParams structure
type GetCustomerOrdersByCustomerIDParams struct {
	CustomerID string
}

// GetCustomerOrdersByCustomerID function
func GetCustomerOrdersByCustomerID(params GetCustomerOrdersByCustomerIDParams) ([]CustomerOrder, error) {
	rows, err := db.Query("SELECT id, customer_id, total_price, created_at, updated_at FROM customer_orders WHERE customer_id = $1", params.CustomerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []CustomerOrder
	for rows.Next() {
		var order CustomerOrder
		err := rows.Scan(&order.ID, &order.CustomerID, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
