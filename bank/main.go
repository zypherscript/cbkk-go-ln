package main

import (
	"bank/repository"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	connStr := "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable"
	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	customerRepository := repository.NewCustomerRepositoryDb(db)
	customers, err := customerRepository.GetAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, customer := range customers {
		fmt.Println(customer)
	}

	customer, err := customerRepository.GetById(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(*customer)
}
