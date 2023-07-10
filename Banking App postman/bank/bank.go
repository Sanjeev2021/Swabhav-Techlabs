package bank

import (
	//account "BankingApp/Account"
	"errors"

	uuid "github.com/satori/go.uuid"

	"bankingApp/account"
)

type Bank struct {
	ID       uuid.UUID
	BankName string
	PassBook []string
	Accounts []*account.Account
}

var BanksCreatedByMe []*Bank

func CreateBank(BankName string) *Bank {
	newBank := &Bank{
		ID:       uuid.NewV4(),
		BankName: BankName,
	}
	BanksCreatedByMe = append(BanksCreatedByMe, newBank)
	return newBank
}

func FindBank(BankSlice []*Bank, id string) (*Bank, bool) {
	for i := 0; i < len(BankSlice); i++ {
		if BankSlice[i].ID.String() == id {
			return BankSlice[i], true
		}
	}
	return nil, false
}

func FindBankByName(bankSlice []*Bank, bankname string) (*Bank, bool) {
	for i := 0; i < len(bankSlice); i++ {
		if bankSlice[i].BankName == bankname {
			return bankSlice[i], true
		}
	}
	return nil, false
}

func DeleteBank(id string) (*Bank, error) {
	for i, Bank := range BanksCreatedByMe {
		if Bank.ID.String() == id {
			BanksCreatedByMe = append(BanksCreatedByMe[:i], BanksCreatedByMe[i+1:]...)
			return Bank, nil
		}
	}
	return nil, errors.New("no Bank found")
}

func UpdateBank(id, field, value string) (*Bank, error) {
	BankToUpdate, BankExist := FindBank(BanksCreatedByMe, id)
	if !BankExist {
		return nil, errors.New("Bank does not exist")
	}
	switch field {
	case "BankName":
		BankToUpdate.BankName = value
	}
	return BankToUpdate, nil
}

func AddAccount(bank *Bank, account *account.Account) {
	bank.Accounts = append(bank.Accounts, account)
}
