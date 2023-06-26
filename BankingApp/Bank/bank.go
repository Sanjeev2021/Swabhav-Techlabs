package bank

import (
	//account "BankingApp/Account"
	//"errors"

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

func NewBank(bankname string) *Bank {
	return &Bank{
		ID:       uuid.NewV4(),
		BankName: bankname,
		IsAdmin:  true,
		Account:  &Account.Account{},
	}
}

func FindBank(bankSlice []Bank, bankname string) (*Bank, bool) {
	for i := 0; i < len(bankSlice); i++ {
		if bankSlice[i].Bankname == bankname {
			return &bankSlice[i], true
		}
	}
	return nil, false
}

func CreateNewAccount(bankname string, accountbalance float64, passbook string) (*Account.Account, error) {
	// if !IsAdmin {
	// 	return nil, errors.New("YOU ARE ADMIN YOU CANT CREATE ACCOUNT")
	// }

	createAccount, _ := Account.NewAccount(bankname, accountbalance, passbook)
	return createAccount, nil
}

func (b *Bank) DeleteBank(bankname string) error {
	for i := 0; i < len(b.BanksCreatedByMe); i++ {
		if b.BanksCreatedByMe[i].Bankname == bankname {
			b.BanksCreatedByMe = append(b.BanksCreatedByMe[:i], b.BanksCreatedByMe[i+1:]...)
			return nil
		}

	}
	return nil
}

func (b *Bank) UpdateBank(bankname string) (*Bank, error) {
	bankToUpdate, found := FindBank(b.BanksCreatedByMe, bankname)
	if found {
		bankToUpdate.Bankname = bankname
		return bankToUpdate, nil
	}
	return nil, nil
}
