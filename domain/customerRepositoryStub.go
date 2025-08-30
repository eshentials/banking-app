package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// func NewCustomerRepositoryStub() CustomerRepository {
// 	return CustomerRepositoryStub{
// 		customers: []Customer{
// 			{ID: "1", Name: "John", City: "New York", Zipcode: 10001, DateOfBirth: "1990-01-01", Status: "1"},
// 			{ID: "2", Name: "Jane", City: "Los Angeles", Zipcode: 90001, DateOfBirth: "1985-05-15", Status: "1"},
// 			{ID: "3", Name: "Mike", City: "Chicago", Zipcode: 60601, DateOfBirth: "1978-09-23", Status: "0"},
// 		},
// 	}
// }
