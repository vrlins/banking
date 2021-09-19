package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vrlins/banking/dto"
	"github.com/vrlins/banking/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (handler *AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err == nil {
		request.CustomerId = customerId
		account, appError := handler.service.NewAccount(request)

		if appError == nil {
			writeResponse(w, http.StatusCreated, account)
		} else {
			writeResponse(w, appError.Code, appError.AsMessage())
		}
	} else {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}
}

// /customers/2000/accounts/90720
func (h AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		//build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := h.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}

}
