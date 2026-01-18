package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type customerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) customerHandler {
	return customerHandler{customerService: customerService}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handleError(w, errs.NewMethodNotAllowedError())
		return
	}
	ctx := r.Context()
	customers, err := h.customerService.GetCustomers(ctx)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request, customerID int) {
	if r.Method != http.MethodGet {
		handleError(w, errs.NewMethodNotAllowedError())
		return
	}
	ctx := r.Context()

	customer, err := h.customerService.GetCustomer(ctx, customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (h customerHandler) HandleCustomer(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/customers" {
		h.GetCustomers(w, r)
		return
	}
	parts := strings.Split(path, "/")
	if len(parts) == 3 && parts[1] == "customers" {
		customerID, err := strconv.Atoi(parts[2])
		if err != nil {
			handleError(w, errs.NewBadRequestError())
			return
		}
		h.GetCustomer(w, r, customerID)
		return
	}
	http.NotFound(w, r)
}
