package service

import (
	"context"
	"time"
)

type NewAccountRequest struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type AccountResponse struct {
	AccountID   int       `json:"account_id"`
	OpeningDate time.Time `json:"opening_date"`
	AccountType string    `json:"account_type"`
	Amount      float64   `json:"amount"`
	Status      int       `json:"status"`
}

type AccountService interface {
	Create(ctx context.Context, customerID int, account NewAccountRequest) (*AccountResponse, error)
	GetAll(ctx context.Context, customerID int) ([]AccountResponse, error)
}
