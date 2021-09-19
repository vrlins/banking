package domain

import (
	"github.com/vrlins/banking/dto"
	"github.com/vrlins/banking/errs"
)

type FilterByStatus int

const (
	Inactive FilterByStatus = iota
	Active
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status FilterByStatus) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}

func NewFilterByStatus(status string) FilterByStatus {
	switch status {
	case "active":
		return Active
	case "inactive":
		return Inactive
	}
	return -1
}
