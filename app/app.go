package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/vrlins/banking/domain"
	"github.com/vrlins/banking/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined.")
	}
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:abcd1234@tcp(localhost:3306)/banking")

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

func Start() {
	sanityCheck()

	dbClient := getDbClient()

	customerRepository := domain.NewCustomerRepositoryDb(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepository)}

	accountRepository := domain.NewAccountRepositoryDb(dbClient)

	ah := AccountHandlers{service.NewAccountService(accountRepository)}

	router := mux.NewRouter()
	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet)

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet)

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).
		Methods(http.MethodPost)

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
