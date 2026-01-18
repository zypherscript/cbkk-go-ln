package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type accountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) accountHandler {
	return accountHandler{accountService: accountService}
}

func (h accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request, customerID int) {
	ctx := r.Context()

	if r.Header.Get("Content-Type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	customers, err := h.accountService.Create(ctx, customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request, customerID int) {
	ctx := r.Context()

	account, err := h.accountService.GetAll(ctx, customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (h accountHandler) HandleAccount(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) == 4 && parts[1] == "customers" && parts[3] == "accounts" {
		customerID, err := strconv.Atoi(parts[2])
		if err != nil {
			handleError(w, errs.NewBadRequestError())
			return
		}
		if r.Method == http.MethodGet {
			h.GetAccounts(w, r, customerID)
		} else if r.Method == http.MethodPost {
			h.CreateAccount(w, r, customerID)
		} else {
			handleError(w, errs.NewMethodNotAllowedError())
		}
		return
	}
	http.NotFound(w, r)
}
