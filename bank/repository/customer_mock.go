package repository

import (
	"context"
	"errors"
	"time"
)

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 101, Name: "Alice Johnson", DateOfBirth: time.Date(1990, time.March, 15, 0, 0, 0, 0, time.UTC), City: "Bangkok", ZipCode: "10110", Status: 1},
	}
	return customerRepositoryMock{customers: customers}
}

func (m customerRepositoryMock) GetAll(ctx context.Context) ([]Customer, error) {
	return m.customers, nil
}

func (m customerRepositoryMock) GetById(ctx context.Context, id int) (*Customer, error) {
	for i, v := range m.customers {
		if v.CustomerID == id {
			return &m.customers[i], nil
		}
	}
	return nil, errors.New("customer not found")
}
