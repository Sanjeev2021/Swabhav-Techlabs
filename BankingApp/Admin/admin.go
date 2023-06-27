package admin

import (
	"errors"

	uuid "github.com/satori/go.uuid"

	bank "bankingapp/Bank"
	user "bankingapp/User"
)

type Admin struct {
	ID               uuid.UUID
	AdminName        string
	BanksCreatedByMe []*bank.Bank
	IsAdmin          bool
	UsersCreatedByMe []*user.User
}

// Creating a new admin
func NewAdmin(adminname string) *Admin {
	return &Admin{
		ID:               uuid.NewV4(),
		AdminName:        adminname,
		IsAdmin:          true,
		BanksCreatedByMe: make([]*bank.Bank, 0),
		UsersCreatedByMe: make([]*user.User, 0),
	}
}

// Allowing admin to create new users
func (ad *Admin) CreateUser(firstname, lastname string) (*user.User, error) {
	// Validations
	if !ad.IsAdmin {
		return nil, errors.New("not admin")
	}
	// Creating user
	UserCreated := user.NewUser(firstname, lastname, 1000)
	// Appending user to the slice
	ad.UsersCreatedByMe = append(ad.UsersCreatedByMe, UserCreated)
	return UserCreated, nil
}

// Letting admins create banks
func (ad *Admin) CreateBank(bankname string) (*bank.Bank, error) {
	// Validations
	if !ad.IsAdmin {
		return nil, errors.New("not admin")
	}
	// Creating the new bank
	BankCreated := bank.NewBank(bankname)
	// Appending the bank to the slice
	ad.BanksCreatedByMe = append(ad.BanksCreatedByMe, BankCreated)
	return BankCreated, nil
}

// Deleting a user
func (ad *Admin) DeleteCreatedUser(firstname string) error {
	// Validations
	if !ad.IsAdmin {
		return errors.New("not admin")
	}
	// Looping through all the users
	for i := 0; i < len(ad.UsersCreatedByMe); i++ {
		// If user with the same firstname is found
		if ad.UsersCreatedByMe[i].FirstName == firstname {
			// Delete the user
			ad.UsersCreatedByMe = append(ad.UsersCreatedByMe[:i], ad.UsersCreatedByMe[i+1:]...)
			return nil
		}
	}
	return nil
}

// Updating an already existing user
func (ad *Admin) UpdateCreatedUser(firstname, field, value string) error {
	// Validations
	if !ad.IsAdmin {
		return errors.New("not admin")
	}
	userToUpdate, isUserExist := FindUser(ad.UsersCreatedByMe, firstname)
	if !isUserExist {
		return errors.New("USER DOES NOT EXIST")
	}
	// Updating values
	switch field {
	case "firstName":
		userToUpdate.FirstName = value
	case "lastName":
		userToUpdate.LastName = value
	default:
		return errors.New("not valid field")
	}
	return nil
}

// Finding a user
func FindUser(userSlice []*user.User, firstname string) (*user.User, bool) {
	// Looping through all the users
	for i := 0; i < len(userSlice); i++ {
		// Returning if the user with the same first name of target is found
		if userSlice[i].FirstName == firstname {
			return userSlice[i], true
		}
	}
	return nil, false
}

// Deleting an already existing bank
func (ad *Admin) DeleteCreatedBank(bankname string) error {
	// Validations
	if !ad.IsAdmin {
		return errors.New("not admin")
	}
	// Looping through all the banks
	for i := 0; i < len(ad.BanksCreatedByMe); i++ {
		// Checking if the bank is found
		if ad.BanksCreatedByMe[i].BankName == bankname {
			// Adding the bank to the slice
			ad.BanksCreatedByMe = append(ad.BanksCreatedByMe[:i], ad.BanksCreatedByMe[i+1:]...)
			return nil
		}
	}
	return nil
}

// Updating an already existing bank
func (ad *Admin) UpdateCreatedBank(bankname, field, value string) error {
	// Validations
	if !ad.IsAdmin {
		return errors.New("not admin")
	}
	bankToUpdate, isBankExist := FindBank(ad.BanksCreatedByMe, bankname)
	if !isBankExist {
		return errors.New("BANK DOES NOT EXIST")
	}
	// Updating fields
	switch field {
	case "bankname":
		bankToUpdate.BankName = value
	default:
		return errors.New("not valid field")
	}
	return nil
}

// Finding a bank
func FindBank(bankSlice []*bank.Bank, bankname string) (*bank.Bank, bool) {
	// Looping through all the banks
	for i := 0; i < len(bankSlice); i++ {
		// Returning if the bank with the same name of target is found
		if bankSlice[i].BankName == bankname {
			return bankSlice[i], true
		}
	}
	return nil, false
}
