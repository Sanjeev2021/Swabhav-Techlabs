package Account

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	// bank "bankingapp/Bank"
)

type Account struct {
	ID                uuid.UUID
	PassBook          string
	BankName          string
	AccountCratedByMe []Account
	AccountBalance    float64
	IsUser            bool
	// Transfer          float64
}

func NewAccount(bankname string, accountbalance float64, passbook string) (*Account, error) {
	if accountbalance <= 1000 {
		return nil, errors.New("Insufficient balance")
	}
	return &Account{
		ID:             uuid.NewV4(),
		BankName:       bankname,
		PassBook:       passbook,
		AccountBalance: accountbalance,
		IsUser:         true,
	}, nil
}

func (a *Account) DeleteAccount(bankname string) (*Account, error) {
	for i := 0; i < len(a.AccountCratedByMe); i++ {
		if a.AccountCratedByMe[i].BankName == bankname {
			a.AccountCratedByMe = append(a.AccountCratedByMe[:i], a.AccountCratedByMe[i+1:]...)
			return a, nil
		}

	}
	return nil, nil
}

func (a *Account) UpdateAccount(bankname string) (*Account, error) {
	AccountToUpdate, found := FindAccount(a.AccountCratedByMe, bankname)
	if !found {
		return nil, nil
	}
	return AccountToUpdate, nil

}

func FindAccount(accountSlice []Account, bankname string) (*Account, bool) {
	for i := 0; i < len(accountSlice); i++ {
		if accountSlice[i].BankName == bankname {
			return &accountSlice[i], true
		}
	}
	return nil, false
}

func (a *Account) Deposit(amount float64) {
	a.AccountBalance += amount
}

func (a *Account) Withdraw(amount float64) {
	a.AccountBalance -= amount
}

func (a *Account) TransferMoney(amount float64, account *Account) error {
	if a.AccountBalance < amount {
		return errors.New("Insufficient balance")
	}
	a.AccountBalance -= amount
	account.AccountBalance += amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.AccountBalance
}
