package domain

import (
	"errors"
	"strings"
)

type AccountType string

const (
	AccountTypeSaving   AccountType = "saving"
	AccountTypeChecking AccountType = "checking"
)

func ValidateAccountType(accountType string) error {
	at := AccountType(strings.ToLower(accountType))

	switch at {
	case AccountTypeSaving, AccountTypeChecking:
		return nil
	default:
		return errors.New("account type mismatch")
	}
}
