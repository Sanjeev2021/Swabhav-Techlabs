package bank

import (
	uuid "github.com/satori/go.uuid"

	"bankingapp/Account"
)

type Bank struct {
	ID               uuid.UUID
	BankName         string
	Account          *Account.Account
	IsAdmin          bool
	BanksCreatedByMe []Bank
	Bankname         string
}

// Creating a new bank
func NewBank(bankname string) *Bank {
	return &Bank{
		ID:       uuid.NewV4(),
		BankName: bankname,
		IsAdmin:  true,
		Account:  &Account.Account{},
	}
}

// Finding if a bank exists
func FindBank(bankSlice []Bank, bankname string) (*Bank, bool) {
	// Looping through all the banks
	for i := 0; i < len(bankSlice); i++ {
		// Checking if the bank exists and returning if it does
		if bankSlice[i].Bankname == bankname {
			return &bankSlice[i], true
		}
	}
	return nil, false
}

// Creating a new account in the respective bank
func CreateNewAccount(bankname string, accountbalance float64, passBook *[]string) (*Account.Account, error) {
	// Creating a new account
	createAccount, err := Account.NewAccount(bankname, accountbalance, passBook)
	if err != nil {
		return nil, err
	}
	return createAccount, nil
}

// Deleting the bank
func (b *Bank) DeleteBank(bankname string) error {
	// Looping through all the banks
	for i := 0; i < len(b.BanksCreatedByMe); i++ {
		// Checking if the bank exists and deleting it if it does
		if b.BanksCreatedByMe[i].Bankname == bankname {
			b.BanksCreatedByMe = append(b.BanksCreatedByMe[:i], b.BanksCreatedByMe[i+1:]...)
			return nil
		}

	}
	return nil
}

// Updating an already existing bank
func (b *Bank) UpdateBank(bankname string) (*Bank, error) {
	// Finding the bank to update
	bankToUpdate, found := FindBank(b.BanksCreatedByMe, bankname)
	if found {
		bankToUpdate.Bankname = bankname
		return bankToUpdate, nil
	}
	return nil, nil
}
