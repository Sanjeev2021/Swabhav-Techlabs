package admin

import (
	"errors"

	uuid "github.com/satori/go.uuid"

	//"bankingapp/Account"
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

func NewAdmin(adminname string) *Admin {
	return &Admin{
		ID:               uuid.NewV4(),
		AdminName:        adminname,
		IsAdmin:          true,
		BanksCreatedByMe: make([]*bank.Bank, 0),
		UsersCreatedByMe: make([]*user.User, 0),
	}
}

func (ad *Admin) CreateUser(firstname, lastname string) (*user.User, error) {
	if !ad.IsAdmin {
		return nil, errors.New("Not admin")
	}

	UserCreated := user.NewUser(firstname, lastname, 1000)

	ad.UsersCreatedByMe = append(ad.UsersCreatedByMe, UserCreated)
	return UserCreated, nil

}

func (ad *Admin) CreateBank(bankname string) (*bank.Bank, error) {
	if !ad.IsAdmin {
		return nil, errors.New("not admin")
	}

	BankCreated := bank.NewBank(bankname)
	ad.BanksCreatedByMe = append(ad.BanksCreatedByMe, BankCreated)
	return BankCreated, nil
}

func (ad *Admin) DeleteCreatedUser(firstname string) error {
	if !ad.IsAdmin {
		return errors.New("not admin")
	}
	for i := 0; i < len(ad.UsersCreatedByMe); i++ {
		if ad.UsersCreatedByMe[i].FirstName == firstname {
			ad.UsersCreatedByMe = append(ad.UsersCreatedByMe[:i], ad.UsersCreatedByMe[i+1:]...)
			return nil
		}
	}
	return nil
}

func (ad *Admin) UpdateCreatedUser(firstname, field, value string) error {
	if !ad.IsAdmin {
		return errors.New("not admin")
	}
	userToUpdate, isUserExist := FindUser(ad.UsersCreatedByMe, firstname)
	if !isUserExist {
		return errors.New("USER DOES NOT EXIST")
	}
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

func FindUser(userSlice []*user.User, firstname string) (*user.User, bool) {
	for i := 0; i < len(userSlice); i++ {
		if userSlice[i].FirstName == firstname {
			return userSlice[i], true
		}
	}
	return nil, false
}

func (ad *Admin) DeleteCreatedBank(bankname string) error {
	if !ad.IsAdmin {
		return errors.New("not admin")
	}
	for i := 0; i < len(ad.BanksCreatedByMe); i++ {
		if ad.BanksCreatedByMe[i].BankName == bankname {
			ad.BanksCreatedByMe = append(ad.BanksCreatedByMe[:i], ad.BanksCreatedByMe[i+1:]...)
			return nil
		}
	}
	return nil
}

func (ad *Admin) UpdateCreatedBank(bankname, field, value string) error {
	if !ad.IsAdmin {
		return errors.New("not admin")
	}
	bankToUpdate, isBankExist := FindBank(ad.BanksCreatedByMe, bankname)
	if !isBankExist {
		return errors.New("BANK DOES NOT EXIST")
	}
	switch field {
	case "bankname":
		bankToUpdate.BankName = value
	default:
		return errors.New("not valid field")
	}
	return nil
}

// func (ad *Admin) FindBank(bankname string) (*bank.Bank, bool) {
// 	for i := 0; i < len(ad.BanksCreatedByMe); i++ {
// 		if ad.BanksCreatedByMe[i].BankName == bankname {
// 			return &ad.BanksCreatedByMe[i], true
// 		}
// 	}
// 	return nil, false
// }

func FindBank(bankSlice []*bank.Bank, bankname string) (*bank.Bank, bool) {
	for i := 0; i < len(bankSlice); i++ {
		if bankSlice[i].BankName == bankname {
			return bankSlice[i], true
		}
	}
	return nil, false
}
