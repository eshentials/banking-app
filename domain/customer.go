package domain

import "github.com/eshentials/banking-app/errs"

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     int    `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	//status ==1 status==0 status=="", return all
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
