package dto

import (
	"strings"

	"github.com/vrlins/banking-lib/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"ammount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}

	accounType := strings.ToLower(r.AccountType)

	if accounType != "saving" && accounType != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}

	return nil
}
