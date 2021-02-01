package domain

import "github.com/PrasadR287/banking/errs"

// Customer dto
type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// CustomerRepository interface
type CustomerRepository interface {
	// status == 1 status ==0 status ==""
	FindAll(status string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
