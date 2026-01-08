package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type customerRepositoryDB struct {
	db *pgxpool.Pool
}

func NewCustomerRepositoryDb(db *pgxpool.Pool) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll(ctx context.Context) ([]Customer, error) {
	err := r.db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []Customer{}
	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.CustomerID, &customer.Name, &customer.DateOfBirth, &customer.City, &customer.ZipCode, &customer.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(ctx context.Context, id int) (*Customer, error) {
	err := r.db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = $1"
	var customer Customer
	err = r.db.QueryRow(ctx, query, id).Scan(&customer.CustomerID, &customer.Name, &customer.DateOfBirth, &customer.City, &customer.ZipCode, &customer.Status)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
