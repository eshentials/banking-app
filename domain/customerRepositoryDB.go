package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eshentials/banking-app/errs"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customer := Customer{}
	err := d.client.QueryRow("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?", id).
		Scan(&customer.ID, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &customer, nil
}

// Factory function to create a new repository
func NewCustomerRepositoryDB() (*CustomerRepositoryDB, error) {
	// Read env vars
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	client, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	}

	// Connection pool settings
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &CustomerRepositoryDB{client: client}, nil
}

func (repo *CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status != "" {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		rows, err = repo.client.Query(findAllSQL)
	} else {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status=?"
		rows, err = repo.client.Query(findAllSQL, status)
	}
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	defer rows.Close()

	customers := []Customer{}
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}
