package account

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"

	"bankingApp/user"
)

type Account struct {
	ID             uuid.UUID
	BankName       string
	AccountType    string
	PassBook       []string
	AccountBalance float64
	Owner          *user.User
}

var AccountsCreatedByMe []*Account

func CreateAccount(bankname, accounttype string, accountbalance float64, owner *user.User) (*Account, error) {
	if accountbalance < 1000 {
		return nil, errors.New("minimum of $1000 needed to create account")
	}
	newAccount := &Account{
		ID:             uuid.NewV4(),
		BankName:       bankname,
		AccountType:    accounttype,
		PassBook:       []string{},
		AccountBalance: accountbalance,
		Owner:          owner,
	}
	AccountsCreatedByMe = append(AccountsCreatedByMe, newAccount)
	return newAccount, nil
}

func UpdateAccount(id, field, value string) (*Account, error) {
	AccountToUpdate, AccountExist := FindAccount(AccountsCreatedByMe, id)
	if !AccountExist {
		return nil, errors.New("Account does not exist")
	}
	switch field {
	case "BankName":
		AccountToUpdate.BankName = value
	case "AccountType":
		AccountToUpdate.AccountType = value
	}
	return AccountToUpdate, nil
}

func FindAccount(slice []*Account, id string) (*Account, bool) {
	for i := 0; i < len(slice); i++ {
		if slice[i].ID.String() == id {
			return slice[i], true
		}
	}
	return nil, false
}

func DeleteAccount(id string) (*Account, error) {
	// index out of range so used range instead of for
	for i, account := range AccountsCreatedByMe {
		if account.ID.String() == id {
			AccountsCreatedByMe = append(AccountsCreatedByMe[:i], AccountsCreatedByMe[i+1:]...)
			return account, nil
		}
	}
	return nil, errors.New("no acccount found")
}

func (a *Account) WithdrawMoney(amount float64) error {
	if amount > a.AccountBalance {
		return errors.New("insufficient funds")
	}
	a.AccountBalance -= amount
	return nil
}

func (a *Account) DepositMoney(amount float64) {
	a.AccountBalance += amount
}

func GetCurrentTime() string {
	dt := time.Now()
	return dt.Format("01-02-2006 15:04:05")
}
