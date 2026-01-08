package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	r := mux.NewRouter()
	r.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods("GET")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
