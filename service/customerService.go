package service

import "github.com/eshentials/banking-app/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(r domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: r}
}
