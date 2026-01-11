package handler

import (
	"bank/errs"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		http.Error(w, e.Message, e.Code)
	case error:
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
