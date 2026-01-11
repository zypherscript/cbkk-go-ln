package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) customerService {
	return customerService{customerRepository: customerRepository}
}

func (s customerService) GetCustomers(ctx context.Context) ([]CustomerResponse, error) {
	customers, err := s.customerRepository.GetAll(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customerResponses := []CustomerResponse{}
	for _, v := range customers {
		customerResponse := CustomerResponse{v.CustomerID, v.Name, v.Status}
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses, nil
}

func (s customerService) GetCustomer(ctx context.Context, id int) (*CustomerResponse, error) {
	customer, err := s.customerRepository.GetById(ctx, id)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NewNotfoundError("customer not found")
		}
		return nil, errs.NewUnexpectedError()
	}
	customerResponse := CustomerResponse{customer.CustomerID, customer.Name, customer.Status}
	return &customerResponse, nil
}
