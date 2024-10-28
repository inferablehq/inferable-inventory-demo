package main

import (
	"os"

	inferable "github.com/inferablehq/inferable/sdk-go"
)

func registerInferableFunctions() error {
	// Initialize Inferable client
	client, err := inferable.New(inferable.InferableOptions{
		APISecret: os.Getenv("INFERABLE_API_SECRET"),
	})
	if err != nil {
		return err
	}

	// Register Customer Service
	customerService, err := client.RegisterService("CustomerService")
	if err != nil {
		return err
	}

	// Register Customer functions
	_, err = customerService.RegisterFunc(inferable.Function{
		Func:        CreateCustomer,
		Name:        "CreateCustomer",
		Description: "Creates a new customer with the given name and email",
	})
	if err != nil {
		return err
	}

	_, err = customerService.RegisterFunc(inferable.Function{
		Func:        GetCustomer,
		Name:        "GetCustomer",
		Description: "Retrieves a customer by their ID",
	})
	if err != nil {
		return err
	}

	_, err = customerService.RegisterFunc(inferable.Function{
		Func:        GetCustomerByEmail,
		Name:        "GetCustomerByEmail",
		Description: "Retrieves a customer by their email",
	})
	if err != nil {
		return err
	}

	_, err = customerService.RegisterFunc(inferable.Function{
		Func:        UpdateCustomer,
		Name:        "UpdateCustomer",
		Description: "Updates a customer's information",
	})
	if err != nil {
		return err
	}

	_, err = customerService.RegisterFunc(inferable.Function{
		Func:        DeleteCustomer,
		Name:        "DeleteCustomer",
		Description: "Deletes a customer by their ID",
	})
	if err != nil {
		return err
	}

	_, err = customerService.RegisterFunc(inferable.Function{
		Func:        ListCustomers,
		Name:        "ListCustomers",
		Description: "Lists all customers",
	})
	if err != nil {
		return err
	}

	// Register Order Service
	orderService, err := client.RegisterService("OrderService")
	if err != nil {
		return err
	}

	// Register Order functions
	_, err = orderService.RegisterFunc(inferable.Function{
		Func:        CreateCustomerOrder,
		Name:        "CreateOrder",
		Description: "Creates a new order for a customer",
	})
	if err != nil {
		return err
	}

	_, err = orderService.RegisterFunc(inferable.Function{
		Func:        GetCustomerOrder,
		Name:        "GetOrder",
		Description: "Retrieves an order by its ID",
	})
	if err != nil {
		return err
	}

	_, err = orderService.RegisterFunc(inferable.Function{
		Func:        UpdateCustomerOrder,
		Name:        "UpdateOrder",
		Description: "Updates an existing order",
	})
	if err != nil {
		return err
	}

	_, err = orderService.RegisterFunc(inferable.Function{
		Func:        DeleteCustomerOrder,
		Name:        "DeleteOrder",
		Description: "Deletes an order by its ID",
	})
	if err != nil {
		return err
	}

	_, err = orderService.RegisterFunc(inferable.Function{
		Func:        ListCustomerOrders,
		Name:        "ListOrders",
		Description: "Lists all orders",
	})
	if err != nil {
		return err
	}

	_, err = orderService.RegisterFunc(inferable.Function{
		Func:        GetCustomerOrdersByCustomerID,
		Name:        "GetOrdersByCustomer",
		Description: "Lists all orders for a specific customer",
	})
	if err != nil {
		return err
	}

	// Start both services
	err = customerService.Start()
	if err != nil {
		return err
	}

	err = orderService.Start()
	if err != nil {
		return err
	}

	return nil
}
