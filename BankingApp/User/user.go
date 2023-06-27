package user

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"

	"bankingapp/Account"
	bank "bankingapp/Bank"
)

type User struct {
	ID               uuid.UUID
	FirstName        string
	LastName         string
	PassBook         []string
	AccountBalance   *Account.Account
	IsAdmin          bool
	UsersCreatedByMe []*User
	Usersname        string
	Account          *Account.Account
}

// Creating a new user
func NewUser(firstname, lastname string, initialBalance float64) *User {
	user := User{
		ID:             uuid.NewV4(),
		FirstName:      firstname,
		LastName:       lastname,
		AccountBalance: &Account.Account{AccountBalance: initialBalance},
		IsAdmin:        true,
		Account:        &Account.Account{},
		PassBook:       []string{},
	}
	return &user
}

// Finding a user
func FindUser(userSlice []User, usersname string) (*User, bool) {
	// Looping through the slice
	for i := 0; i < len(userSlice); i++ {
		// Checking if the user exists and if it does returning it
		if userSlice[i].Usersname == usersname {
			return &userSlice[i], true
		}
	}
	return nil, false
}

// Creating a new admin
func NewAdmin(firstName, lastName, usersname string) *User {
	return &User{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		Usersname: usersname,
		IsAdmin:   true,
	}
}

// Creating a new account
func (u *User) CreateNewAccount(bankname string, accountbalance float64) (*Account.Account, error) {
	// Validations
	if !u.IsAdmin {
		return nil, errors.New("YOU ARE NOT AN ADMIN YOU CANT CREATE ACCOUNT")
	}
	// Creating the account
	createAccount, _ := bank.CreateNewAccount(bankname, accountbalance, &u.PassBook)

	return createAccount, nil
}

// Finding if an account exits
func (u *User) FindAccount(bankname string) (*Account.Account, bool) {
	// Looping through all the accounts
	for i := 0; i < len(u.UsersCreatedByMe); i++ {
		// Checking if the account exists and returning it if it does
		if u.UsersCreatedByMe[i].Account.BankName == bankname {
			return u.UsersCreatedByMe[i].Account, true
		}
	}
	return nil, false
}

// Deleting an already existing user
func (u *User) DeleteUser(firstname string) error {
	// Looping through all the users
	for i := 0; i < len(u.UsersCreatedByMe); i++ {
		// Checking if the user exists and deleting it if it does
		if u.UsersCreatedByMe[i].FirstName == firstname {
			u.UsersCreatedByMe = append(u.UsersCreatedByMe[:i], u.UsersCreatedByMe[i+1:]...)
			return nil
		}
	}
	return errors.New("no user found")
}

// Updating an already existing user
func (u *User) UpdateUser(firstname string, UsersCreatedByMe []User) error {
	// Finding the user
	userToUpdate, userExist := FindUser(UsersCreatedByMe, firstname)
	if !userExist {
		return errors.New("user does not exist")
	}
	userToUpdate.FirstName = "NewName"

	return nil
}

// Deleting an already exisiting account
func (u *User) DeleteAccount(bankname string) (*Account.Account, error) {
	// Looping through the slice
	for i := 0; i < len(u.UsersCreatedByMe); i++ {
		// Checking if the account exists and deleting it if it does
		if u.UsersCreatedByMe[i].FirstName == bankname {
			u.UsersCreatedByMe = append(u.UsersCreatedByMe[:i], u.UsersCreatedByMe[i+1:]...)
			return u.AccountBalance, nil
		}

	}
	return nil, nil
}

// Updating an already existing account
func (u *User) UpdateAccount(bankname string, UsersCreatedByMe []User) error {
	// Finding the account
	userToUpdate, userExist := FindUser(UsersCreatedByMe, bankname)
	if !userExist {
		return errors.New("user does not exist")
	}
	userToUpdate.FirstName = "NewName"

	return nil
}

// Depositiong money to the account
func (u *User) DepositMoney(bankname string, amount float64) error {
	// Validations
	if !u.IsAdmin {
		return errors.New("YOU ARE NOT AN ADMIN YOU CANT DEPOSIT MONEY")
	}
	// Depositing the money
	u.AccountBalance.Deposit(amount)
	return nil
}

// Withdrawing the money from the account
func (u *User) WithdrawMoney(bankname string, amount float64) error {
	// Validations
	if !u.IsAdmin {
		return errors.New("YOU ARE NOT AN ADMIN YOU CANT WITHDRAW MONEY")
	}
	// Withdrawing the money
	u.AccountBalance.Withdraw(amount)
	return nil
}

// Transferring money from one account to another
func (u *User) TransferMoney(bankname string, amount float64) error {
	// Validations
	if !u.IsAdmin {
		return errors.New("YOU ARE NOT AN ADMIN YOU CANT TRANSFER MONEY")
	}
	// Transferring the money
	transfer := u.Account.TransferMoney(amount, u.Account)
	if transfer == nil {
		return errors.New("transfer failed")
	}
	return nil
}

// Getting the account balance
func (u *User) GetAccountBalance(bankname string) (float64, error) {
	// Validations
	if !u.IsAdmin {
		return 0, errors.New("YOU ARE NOT AN ADMIN YOU CANT GET ACCOUNT BALANCE")
	}
	// Getting the account balance
	account, found := u.FindAccount(bankname)
	if !found {
		return 0, errors.New("account not found")
	}
	return account.GetBalance(), nil
}

// Getting the passbook
func (u *User) GetPassBook() {
	// Looping through the slice
	for i := 0; i < len(u.PassBook); i++ {
		// Printing the passbook
		fmt.Println(u.PassBook[i])
	}
}
