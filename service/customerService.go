package service

import (
	"github.com/PrasadR287/banking/domain"
	"github.com/PrasadR287/banking/errs"
)

// CustomerService interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService class implements
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer interface method
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ByID(id)
}

// NewCustomerService method
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
