package domain

import "github.com/PrasadR287/banking/errs"

// Customer dto
type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// CustomerRepository interface
type CustomerRepository interface {
	// status == 1 status ==0 status ==""
	FindAll(status string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
