package contact

//crud for contact
import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Contact struct {
	ID           uuid.UUID
	ContactName  string
	ContactType  string
	ContactValue interface{}
}

// func NewContact(contactname string,mycontactinfo []contactinfo.ContactInfo) *Contact {
func NewContact(contactName string, contactType string, contactValue interface{}) Contact {
	return Contact{
		ID:           uuid.NewV4(),
		ContactName:  contactName,
		ContactType:  contactType,
		ContactValue: contactValue,
	}
}

func (c *Contact) UpdateContactName(name string) {
	c.ContactName = name
}

func DeleteContact(contacts []*Contact, contact *Contact) ([]*Contact, error) {
	for i := 0; i < len(contacts); i++ {
		if contacts[i] == contact {
			return append(contacts[:i], contacts[i+1:]...), nil
		}

	}
	return nil, errors.New("no contact found")
}

/*
func (c *Contact) CreateContactInfo(cit, civ string) (*Contact, error) {
	newContactInfo := contactinfo.NewContactInfo(cit, civ)
	c.MyContactInfo = append(c.MyContactInfo, *newContactInfo)
	return nil, errors.New("Name already exist")
}
*/

func FindContactById(id string, contactSlice []Contact) (*Contact, error) {
	for i := 0; i < len(contactSlice); i++ {
		if contactSlice[i].ID.String() == id {
			return &contactSlice[i], nil
		}
	}
	return nil, errors.New("no contact found")
}
