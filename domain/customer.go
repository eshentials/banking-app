package domain

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     int    `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
