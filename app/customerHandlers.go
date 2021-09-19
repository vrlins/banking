package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vrlins/banking-lib/errs"
	"github.com/vrlins/banking/domain"
	"github.com/vrlins/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (handler *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	if status != "inactive" && status != "active" && status != "" {
		writeResponse(w, http.StatusBadRequest, errs.NewValidationError("Invalid parameters"))
		return
	}

	customers, err := handler.service.GetAllCustomers(domain.NewFilterByStatus(status))

	if err == nil {
		writeResponse(w, http.StatusOK, customers)
	} else {
		writeResponse(w, http.StatusInternalServerError, err.AsMessage())
	}
}

func (handler *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	customer, err := handler.service.GetCustomer(customerId)

	if err == nil {
		writeResponse(w, http.StatusOK, customer)
	} else {
		writeResponse(w, err.Code, err.AsMessage())
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
