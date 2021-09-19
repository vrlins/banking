package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/vrlins/banking-lib/errs"
	"github.com/vrlins/banking-lib/logger"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (repository CustomerRepositoryDb) FindAll(status FilterByStatus) ([]Customer, *errs.AppError) {
	var err error

	customers := make([]Customer, 0)

	if status == -1 {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = repository.client.Select(&customers, findAllSql)
	} else {
		findAllSqlWithStatus := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = repository.client.Select(&customers, findAllSqlWithStatus, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unable to get customer list.")
	}

	return customers, nil
}

func (repository CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var customer Customer

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	err := repository.client.Get(&customer, findAllSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found.")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &customer, nil
}

func NewCustomerRepositoryDb(client *sqlx.DB) CustomerRepository {
	return CustomerRepositoryDb{client}
}
