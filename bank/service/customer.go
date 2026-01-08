package service

import (
	"context"
)

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CustomerService interface {
	GetCustomers(ctx context.Context) ([]CustomerResponse, error)
	GetCustomer(ctx context.Context, id int) (*CustomerResponse, error)
}
