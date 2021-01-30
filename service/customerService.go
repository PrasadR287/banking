package service

import "github.com/PrasadR287/banking/domain"

// CustomerService interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// DefaultCustomerService class implements
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer interface method
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// NewCustomerService method
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
