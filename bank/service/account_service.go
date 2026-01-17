package service

import (
	"bank/domain"
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"context"
	"time"
)

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return accountService{accountRepository: accountRepository}
}

func (s accountService) Create(ctx context.Context, customerID int, request NewAccountRequest) (*AccountResponse, error) {
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5000")
	}
	if err := domain.ValidateAccountType(request.AccountType); err != nil {
		return nil, errs.NewValidationError(err.Error())
	}

	account := repository.Account{CustomerID: customerID, OpeningDate: time.Now(), AccountType: request.AccountType, Amount: request.Amount, Status: 1}
	newAccount, err := s.accountRepository.Create(ctx, account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	accountResponse := AccountResponse{newAccount.AccountID, newAccount.OpeningDate, newAccount.AccountType, newAccount.Amount, newAccount.Status}
	return &accountResponse, nil
}

func (s accountService) GetAll(ctx context.Context, customerID int) ([]AccountResponse, error) {
	accounts, err := s.accountRepository.GetAll(ctx, customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	accountResponses := []AccountResponse{}
	for _, v := range accounts {
		accountResponse := AccountResponse{v.AccountID, v.OpeningDate, v.AccountType, v.Amount, v.Status}
		accountResponses = append(accountResponses, accountResponse)
	}
	return accountResponses, nil
}
