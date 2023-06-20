package user

// CRUD
import (
	"errors"

	uuid "github.com/satori/go.uuid"

	"contactApp/contact"
)

type User struct {
	ID               uuid.UUID
	FirstName        string
	LastName         string
	IsAdmin          bool
	usersCreatedByMe []User
	Mycontacts       []*contact.Contact
	Username         string
}

func findUser(userSlice []User, username string) (*User, bool) {
	for i := 0; i < len(userSlice); i++ {
		if userSlice[i].Username == username {
			return &userSlice[i], true
		}

	}
	return nil, false
}

func NewAdmin(firstName, lastName, username string) *User {
	return &User{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		IsAdmin:   true,
	}
}

func (u *User) NewUser(firstName, lastName, username string) (*User, error) {
	if !u.IsAdmin {
		return nil, errors.New(u.FirstName + "is not authorized to create a user")
	}
	_, isUserExisit := findUser(u.usersCreatedByMe, username)
	if isUserExisit {
		return nil, errors.New("user already doest not exist")
	}

	newUser := &User{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		IsAdmin:   false,
	}
	u.usersCreatedByMe = append(u.usersCreatedByMe, *newUser)
	return newUser, nil
}

func (u *User) ReadNewUser() ([]User, error) {
	if !u.IsAdmin {
		return nil, errors.New(u.FirstName + "not admin")
	}
	return u.usersCreatedByMe, nil
}

func FindContact(contactSlice []*contact.Contact, contactname string) (*contact.Contact, bool) {
	for i := 0; i < len(contactSlice); i++ {
		if contactSlice[i].ContactName == contactname {
			return contactSlice[i], true
		}
	}
	return nil, false
}

// func FindContactInfo(list []ContactInfo, ID uuid.UUID) (*ContactInfo, error) {
//	for  i := 0; i < len(list); i++ {
//	if list[i].ID == ID {
//		return &list[i], nil
//	}
//	}
//	return nil, errors.New("No contact info found")
//}

// all the user will be updated in this
// func (u *User) UpdateUser() (updatedUser *User, error) {
// 	if !u.IsAdmin {
// 		return errors.New(u.FirstName + "not admin")
// 	}
// 	for i := 0; i < len(u.usersCreatedByMe); i++ {
// 		 u.usersCreatedByMe[i].FirstName = updatedUser.FirstName
// 		if u.usersCreatedByMe[i] = *updatedUser
// 	}
// 	return nil

// }

func (u *User) UpdatedUser(username, field string, value string) (*User, error) {
	if !u.IsAdmin {
		return nil, errors.New(u.FirstName + "not admin")
	}
	userToUpdate, isUserExisit := findUser(u.usersCreatedByMe, username)
	if !isUserExisit {
		return nil, errors.New("user already doest not exist")
	}
	switch field {
	case "firstName":
		userToUpdate.FirstName = value
		return userToUpdate, nil
	case "lastName":
		userToUpdate.LastName = value
		return userToUpdate, nil

	default:
		return nil, errors.New("field is not present")

	}
}

func (u *User) DeleteUser(username string) ([]User, error) {
	if !u.IsAdmin {
		return nil, errors.New(u.FirstName + "not admin")
	}
	_, isUserExist := findUser(u.usersCreatedByMe, username)
	if !isUserExist {
		return nil, errors.New("user already does not exist")
	}
	for i := 0; i < len(u.usersCreatedByMe); i++ {
		if u.usersCreatedByMe[i].Username == username {
			return append(u.usersCreatedByMe[:i], u.usersCreatedByMe[i+1:]...), nil
		}
	}
	return nil, errors.New("User does not exist")
}

// Crud cannot use contact variable of
// CREATE , UPDATE
func (u *User) CreateContactUser(contactname string) (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "not admin")
	}
	_, ContactExisit := FindContact(u.Mycontacts, contactname)
	if ContactExisit {
		return nil, errors.New("contact eixst")
	}
	newcontact := contact.NewContact(contactname)
	u.Mycontacts = append(u.Mycontacts, newcontact)

	return newcontact, nil
}

func (u *User) UpdateContactUser(oldcontactname, contactname string) (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is admin")
	}
	contactToUpdate, contactExist := FindContact(u.Mycontacts, oldcontactname)
	if !contactExist {
		return nil, errors.New("contact does not exist")
	}
	contactToUpdate.UpdateContactName(contactname)
	return contactToUpdate, nil
}

func (u *User) DeleteContactUser(ContactName string) ([]*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is admin")
	}
	_, isUserExist := FindContact(u.Mycontacts, ContactName)
	if !isUserExist {
		return nil, errors.New("contact does not exist")
	}
	for i := 0; i < len(u.Mycontacts); i++ {
		if u.Mycontacts[i].ContactName == ContactName {
			return append(u.Mycontacts[:i], u.Mycontacts[i+1:]...), nil
		}
	}
	return nil, errors.New("User does not exist")

}

// func (u *User) UpdateContactUser( contactname string) (*contact.Contact, error) {
// 	if !u.IsAdmin {
// 		return nil, errors.New(u.FirstName + "not admin")
// 	}
// 	ContactToUpdate, ContactExisit := findUser(u.usersCreatedByMe, contactname)
// 	if !ContactExisit {
// 		return nil, errors.New("Contact does not eixst")
// 	}

// }

func (u *User) CreateContactInfo(contactname, cit, civ string) (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is admin")
	}
	obj, isUserExist := FindContact(u.Mycontacts, contactname)
	if !isUserExist {
		return nil, errors.New("contact does not exist")
	}
	//return obj.CreateContactInfo(cit, civ)
	obj.CreateContactInfo(cit, civ)
	return obj, nil

}

func (u *User) UpdateContactInfo(contactname string, field string, value string) error {
	if u.IsAdmin {
		return errors.New(u.FirstName + "is admin")
	}
	updateCon, contactExist := FindContact(u.Mycontacts, contactname)
	if !contactExist {
		return errors.New("contact does not exist")
	}
	updateCon.UpdateContactInfo(contactname, field, value)

	return nil
}

func (u *User) DeleteContactInfo(contactname, contactInfoType string) (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is admin")
	}
	DeleteContact, contactExist := FindContact(u.Mycontacts, contactname)
	if !contactExist {
		return nil, errors.New("CONTACT DOES NOIT EXIST")
	}
	DeleteContact.DeleteContactInfo(contactInfoType, contactname)
	return DeleteContact, nil
}
