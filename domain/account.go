package domain

import (
	"github.com/PrasadR287/banking/dto"
	"github.com/PrasadR287/banking/errs"
)

// Account domain object
type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountID}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
