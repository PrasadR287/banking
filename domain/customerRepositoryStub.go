package domain

// CustomerRepositoryStub struct stub
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll interface method
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub stub data
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
