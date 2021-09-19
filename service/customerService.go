package service

import (
	"github.com/vrlins/banking/domain"
	"github.com/vrlins/banking/dto"
	"github.com/vrlins/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status domain.FilterByStatus) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers(status domain.FilterByStatus) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := d.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)

	for _, v := range customers {
		response = append(response, v.ToDto())
	}

	return response, nil
}

func (d DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := d.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := customer.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repository}
}
