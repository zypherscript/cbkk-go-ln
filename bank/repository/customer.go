package repository

import (
	"context"
	"time"
)

type Customer struct {
	CustomerID  int       `db:"customer_id"`
	Name        string    `db:"name"`
	DateOfBirth time.Time `db:"date_of_birth"`
	City        string    `db:"city"`
	ZipCode     string    `db:"zipcode"`
	Status      int       `db:"status"`
}

type CustomerRepository interface {
	GetAll(ctx context.Context) ([]Customer, error)
	GetById(ctx context.Context, id int) (*Customer, error)
}
