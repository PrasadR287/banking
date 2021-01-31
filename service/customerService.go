package service

import "github.com/PrasadR287/banking/domain"

// CustomerService interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

// DefaultCustomerService class implements
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer interface method
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.ByID(id)
}

// NewCustomerService method
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
