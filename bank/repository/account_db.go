package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepositoryDB struct {
	db *pgxpool.Pool
}

func NewAccountRepositoryDB(db *pgxpool.Pool) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(ctx context.Context, account Account) (*Account, error) {
	err := r.db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	query := "insert into accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) returning account_id"
	err = r.db.QueryRow(ctx, query, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status).Scan(&account.AccountID)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r accountRepositoryDB) GetAll(ctx context.Context, customerID int) ([]Account, error) {
	err := r.db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id = $1"
	rows, err := r.db.Query(ctx, query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []Account{}
	for rows.Next() {
		account := Account{}
		err = rows.Scan(&account.AccountID, &account.CustomerID, &account.OpeningDate, &account.AccountType, &account.Amount, &account.Status)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
