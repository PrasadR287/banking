package domain

import "github.com/PrasadR287/banking/errs"

// Account domain object
type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	save(Account) (*Account, *errs.AppError)
}
