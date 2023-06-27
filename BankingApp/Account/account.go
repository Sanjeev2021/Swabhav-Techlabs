package Account

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID                uuid.UUID
	PassBook          []string
	BankName          string
	AccountCratedByMe []Account
	AccountBalance    float64
	IsUser            bool
	userPassBook      *[]string
}

// Creating a new account
func NewAccount(bankname string, accountbalance float64, passBook *[]string) (*Account, error) {
	// Validations
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

// Deleting an already existing account
func (a *Account) DeleteAccount(bankname string) (*Account, error) {
	// Looping through all the banks
	for i := 0; i < len(a.AccountCratedByMe); i++ {
		// Checking if the bank exists and deleting it if it does
		if a.AccountCratedByMe[i].BankName == bankname {
			a.AccountCratedByMe = append(a.AccountCratedByMe[:i], a.AccountCratedByMe[i+1:]...)
			return a, nil
		}

	}
	return nil, nil
}

// Updating an already existing account
func (a *Account) UpdateAccount(bankname string) (*Account, error) {
	// Finding the account to update
	AccountToUpdate, found := FindAccount(a.AccountCratedByMe, bankname)
	if !found {
		return nil, nil
	}
	return AccountToUpdate, nil

}

// Findind an account
func FindAccount(accountSlice []Account, bankname string) (*Account, bool) {
	// Looping through all the accounts
	for i := 0; i < len(accountSlice); i++ {
		// Checking if the account exists and returning if it does
		if accountSlice[i].BankName == bankname {
			return &accountSlice[i], true
		}
	}
	return nil, false
}

// Depositing money into the account
func (a *Account) Deposit(amount float64) {
	a.AccountBalance += amount
	// Passbook updation
	date := getCurrentTime()
	a.PassBook = append(a.PassBook, fmt.Sprint(date, " Deposited ", amount))
	*a.userPassBook = append(*a.userPassBook, fmt.Sprint(date, " Deposited ", amount))
}

// Withdraw money from the account
func (a *Account) Withdraw(amount float64) {
	a.AccountBalance -= amount
	// Passbook updation
	date := getCurrentTime()
	a.PassBook = append(a.PassBook, fmt.Sprint(date, " Withdrawn ", amount))
	*a.userPassBook = append(*a.userPassBook, fmt.Sprint(date, " Withdrawn ", amount))
}

// Transfer money from one account to another
func (a *Account) TransferMoney(amount float64, account *Account) error {
	// Validations
	if a.AccountBalance < amount {
		return errors.New("insufficient balance")
	}

	date := getCurrentTime()

	a.AccountBalance -= amount
	// Passbook updation
	a.PassBook = append(a.PassBook, fmt.Sprint(date, " Transfered ", amount, " to ", account.BankName))
	*a.userPassBook = append(*a.userPassBook, fmt.Sprint(date, " Transfered ", amount, " to ", account.BankName))
	account.AccountBalance += amount
	account.PassBook = append(account.PassBook, fmt.Sprint(date, " Received ", amount, " from ", a.BankName))
	*account.userPassBook = append(*account.userPassBook, fmt.Sprint(date, " Received ", amount, " from ", a.BankName))
	return nil
}

// Getting the current user balance
func (a *Account) GetBalance() float64 {
	return a.AccountBalance
}

// Printing the pass book of the account
func (a *Account) GetPassBook() {
	for i := 0; i < len(a.PassBook); i++ {
		fmt.Println(a.PassBook[i])
	}
}

// Getting the current time
func getCurrentTime() string {
	dateTime := time.Now()
	date := dateTime.Format("2006-01-02 17:00:00")
	return date
}
