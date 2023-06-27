package Account

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	// bank "bankingapp/Bank"
)

type Account struct {
	ID                uuid.UUID
	PassBook          []string
	BankName          string
	AccountCratedByMe []Account
	AccountBalance    float64
	IsUser            bool
	userPassBook      *[]string
	// Transfer          float64
}

func NewAccount(bankname string, accountbalance float64, passBook *[]string) (*Account, error) {
	if accountbalance < 1000 {
		return nil, errors.New("intial ammount needs to be atleast 1000")
	}
	return &Account{
		ID:             uuid.NewV4(),
		BankName:       bankname,
		PassBook:       []string{},
		AccountBalance: accountbalance,
		IsUser:         true,
		userPassBook:   passBook,
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
	date := getCurrentTime()
	a.PassBook = append(a.PassBook, fmt.Sprint(date, " Deposited ", amount))
}

func (a *Account) Withdraw(amount float64) {
	a.AccountBalance -= amount
	date := getCurrentTime()
	a.PassBook = append(a.PassBook, fmt.Sprint(date, " Withdrawn ", amount))
}

func (a *Account) TransferMoney(amount float64, account *Account) error {
	if a.AccountBalance < amount {
		return errors.New("insufficient balance")
	}

	date := getCurrentTime()

	a.AccountBalance -= amount
	a.PassBook = append(a.PassBook, fmt.Sprint(date, " Transfered ", amount, " to ", account.BankName))
	*a.userPassBook = append(*a.userPassBook, fmt.Sprint(date, " Transfered ", amount, " to ", account.BankName))
	account.AccountBalance += amount
	account.PassBook = append(account.PassBook, fmt.Sprint(date, " Received ", amount, " from ", a.BankName))
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.AccountBalance
}

func (a *Account) GetPassBook() {
	for i := 0; i < len(a.PassBook); i++ {
		fmt.Println(a.PassBook[i])
	}
}

func getCurrentTime() string {
	dateTime := time.Now()
	date := dateTime.Format("2006-01-02 17:00:00")
	return date
}
