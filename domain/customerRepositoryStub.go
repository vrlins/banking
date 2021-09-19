package domain

import "github.com/vrlins/banking-lib/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (repository CustomerRepositoryStub) FindAll(status FilterByStatus) ([]Customer, *errs.AppError) {
	return repository.customers, nil
}

func (repository CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {
	var customer *Customer

	for i := range repository.customers {
		if repository.customers[i].Id == id {
			customer = &repository.customers[i]
			break
		}
	}

	return customer, nil
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{Id: "1", Name: "Victor", City: "Recife", Zipcode: "1223", DateOfBirth: "1992-01-06", Status: "1"},
		{Id: "2", Name: "Leticia", City: "Recife", Zipcode: "1223", DateOfBirth: "1995-05-21", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
