package account

import (
	"errors"

	"gorm.io/gorm"

	bank "bankingApp/bank"
	"bankingApp/database"
	"bankingApp/services"
)

// Account is a struct that represents a bank account
type Account struct {
	gorm.Model
	ID          uint
	AccountType string
	Balance     float64
	BankID      uint
	Bank        bank.Bank
	UserID      uint
	User        services.User
}

type AccountTranfer struct {
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func init() {
	db := database.GetDB()
	database.Migrate(db, &Account{})
}

// NewAccount creates a new account
func CreateAccount(AccountType string, balance float64, BankId, UserId uint) (*Account, error) {
	db := database.GetDB()

	account := &Account{
		AccountType: AccountType,
		Balance:     balance,
		BankID:      BankId,
		UserID:      UserId,
	}

	result := db.Create(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil

}

func DeleteAccount(id uint) (*Account, error) {
	db := database.GetDB()

	account, err := GetAccountById(id)
	if err != nil {
		return nil, err
	}

	result := db.Delete(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func GetAccountById(id uint) (*Account, error) {
	db := database.GetDB()

	var account Account

	result := db.First(&account, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func UpdateAccount(id uint, updateaccount *Account) (*Account, error) {
	db := database.GetDB()

	account, err := GetAccountById(updateaccount.ID)
	if err != nil {
		return nil, err
	}

	if updateaccount.AccountType != "" {
		account.AccountType = updateaccount.AccountType
	}

	result := db.Save(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func GetAllAccounts(page, pagesize int) ([]*Account, error) {
	db := database.GetDB()

	var accounts []*Account

	result := db.Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}
	return accounts, nil
}

func TransferMoney(FromAccountID, ToAccountID uint, amount float64) error {
	db := database.GetDB()

	fromAccount, err := GetAccountById(FromAccountID)
	if err != nil {
		return err
	}

	toAccount, err := GetAccountById(ToAccountID)
	if err != nil {
		return err
	}

	if fromAccount.Balance < amount {
		return errors.New("insufficient funds")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	result := db.Save(fromAccount)
	if result.Error != nil {
		return result.Error
	}

	result = db.Save(toAccount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DepositMoney(accountId uint, amount float64) error {
	db := database.GetDB()

	account, err := GetAccountById(accountId)
	if err != nil {
		return err
	}

	account.Balance += amount

	result := db.Save(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func WithdrawMoney(accountId uint, amount float64) error {
	db := database.GetDB()

	account, err := GetAccountById(accountId)
	if err != nil {
		return err
	}

	account.Balance -= amount

	if account.Balance < amount {
		return errors.New("insufficient fund")
	}

	result := db.Save(account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
