package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDb()
	defer db.Close()

	customerRepository := repository.NewCustomerRepositoryDb(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryMock
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepository := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	r := mux.NewRouter()
	r.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	r.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.CreateAccount).Methods(http.MethodPost)

	port := viper.GetInt("app.port")
	logs.Info("Server running on :" + strconv.Itoa(port))

	srv := &http.Server{Addr: fmt.Sprintf(":%v", port), Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Error(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logs.Info("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		logs.Error(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDb() *pgxpool.Pool {
	ctx := context.Background()

	connStr := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v",
		viper.GetString("db.driver"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
		viper.GetString("db.sslmode"),
	)
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic(err)
	}
	config.MaxConnLifetime = 3 * time.Minute
	config.MaxConns = 10
	config.MinConns = 5

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(err)
	}

	return db
}
