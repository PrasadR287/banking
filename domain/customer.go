package domain

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
	FindAll() ([]Customer, error)
	ByID(string) (*Customer, error)
}
