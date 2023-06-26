package user

import (
	"errors"
	//"fmt"

	uuid "github.com/satori/go.uuid"

	"bankingapp/Account"
	bank "bankingapp/Bank"

)

type User struct {
	ID               uuid.UUID
	FirstName        string
	LastName         string
	AccountBalance   *Account.Account
	IsAdmin          bool
	UsersCreatedByMe []*User
	Usersname        string
	Account          *Account.Account
}

func NewUser(firstname, lastname string, initialBalance float64) *User {
	return &User{
		ID:             uuid.NewV4(),
		FirstName:      firstname,
		LastName:       lastname,
		AccountBalance: &Account.Account{AccountBalance: initialBalance},
		IsAdmin:        true,
		Account:        &Account.Account{},
	}
}

func FindUser(userSlice []User, usersname string) (*User, bool) {
	for i := 0; i < len(userSlice); i++ {
		if userSlice[i].Usersname == usersname {
			return &userSlice[i], true
		}
	}
	return nil, false
}

func NewAdmin(firstName, lastName, usersname string) *User {
	return &User{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		Usersname: usersname,
		IsAdmin:   true,
	}
}

func (u *User) CreateNewAccount(bankname string, accountbalance float64, passbook string) (*Account.Account, error) {
	if !u.IsAdmin {
		return nil, errors.New("YOU ARE NOT AN ADMIN YOU CANT CREATE ACCOUNT")
	}

	createAccount, _ := bank.CreateNewAccount(bankname, accountbalance, passbook)

	return createAccount, nil
}

func (u *User) FindAccount(bankname string) (*Account.Account, bool) {
	for i := 0; i < len(u.UsersCreatedByMe); i++ {
		if u.UsersCreatedByMe[i].Account.BankName == bankname {
			return u.UsersCreatedByMe[i].Account, true
		}
	}
	return nil, false
}

func (u *User) DeleteUser(firstname string) error {
	for i := 0; i < len(u.UsersCreatedByMe); i++ {
		if u.UsersCreatedByMe[i].FirstName == firstname {
			u.UsersCreatedByMe = append(u.UsersCreatedByMe[:i], u.UsersCreatedByMe[i+1:]...)
			return nil
		}
	}
	return errors.New("no user found")
}

func (u *User) UpdateUser(firstname string, UsersCreatedByMe []User) error {
	userToUpdate, userExist := FindUser(UsersCreatedByMe, firstname)
	if !userExist {
		return errors.New("user does not exist")
	}
	userToUpdate.FirstName = "NewName"

	return nil
}

func (u *User) DeleteAccount(bankname string) (*Account.Account, error) {
	for i := 0; i < len(u.UsersCreatedByMe); i++ {
		if u.UsersCreatedByMe[i].FirstName == bankname {
			u.UsersCreatedByMe = append(u.UsersCreatedByMe[:i], u.UsersCreatedByMe[i+1:]...)
			return u.AccountBalance, nil
		}

	}
	return nil, nil
}

func (u *User) UpdateAccount(bankname string, UsersCreatedByMe []User) error {
	userToUpdate, userExist := FindUser(UsersCreatedByMe, bankname)
	if !userExist {
		return errors.New("user does not exist")
	}
	userToUpdate.FirstName = "NewName"

	return nil
}

func (u *User) DepositMoney(bankname string, amount float64) error {
	if !u.IsAdmin {
		return errors.New("YOU ARE NOT AN ADMIN YOU CANT DEPOSIT MONEY")
	}
	u.AccountBalance.Deposit(amount)
	return nil
}

func (u *User) WithdrawMoney(bankname string, amount float64) error {
	if !u.IsAdmin {
		return errors.New("YOU ARE NOT AN ADMIN YOU CANT WITHDRAW MONEY")
	}
	u.AccountBalance.Withdraw(amount)
	return nil
}

func (u *User) TransferMoney(bankname string, amount float64) error {
	if !u.IsAdmin {
		return errors.New("YOU ARE NOT AN ADMIN YOU CANT TRANSFER MONEY")
	}
	transfer := u.Account.TransferMoney(amount, u.Account)
	if transfer != nil {
		return transfer
	}
	return nil
}

func (u *User) GetAccountBalance(bankname string) (float64, error) {
	if !u.IsAdmin {
		return 0, errors.New("YOU ARE NOT AN ADMIN YOU CANT GET ACCOUNT BALANCE")
	}

	account, found := u.FindAccount(bankname)
	if !found {
		return 0, errors.New("account not found")
	}
	return account.GetBalance(), nil
}
