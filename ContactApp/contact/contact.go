package contact

//crud for contact
import (
	"errors"
	//"fmt"

	uuid "github.com/satori/go.uuid"

	//"contactApp/contact"
	contactinfo "contactApp/contactinfo"
)

type Contact struct {
	ID            uuid.UUID
	ContactName   string
	MyContactInfo []*contactinfo.ContactInfo
}

// func NewContact(contactname string,mycontactinfo []contactinfo.ContactInfo) *Contact {
func NewContact(contactname string) *Contact {
	return &Contact{
		ID:          uuid.NewV4(),
		ContactName: contactname,
		//MyContactInfo: mycontactinfo,
	}

}

func FindContactID(contacts []*Contact, ID uuid.UUID) (*Contact, error) {
	for i := 0; i < len(contacts); i++ {
		if contacts[i].ID == ID {
			return contacts[i], nil
		}
	}
	return nil, errors.New("ID DOESNT EXIST")
	// doubt over there i am using error but down i am doing errors

}

func (c *Contact) UpdateContactName(contactname string) {
	c.ContactName = contactname
}

func DeleteContact(contacts []*Contact, contact *Contact) ([]*Contact, error) {
	for i := 0; i < len(contacts); i++ {
		if contacts[i] == contact {
			return append(contacts[:i], contacts[i+1:]...), nil
		}

	}
	return nil, errors.New("no contact found")
}

func (c *Contact) CreateContactInfo(cit, civ string) (*contactinfo.ContactInfo, error) {
	// _, err := contactinfo.NewContactInfo(c.MyContactInfo, cit, civ)
	// if err != nil {
	// 	return nil, err
	// }
	// for i := 0; i < len(c.MyContactInfo); i++ {
	// 	if c.MyContactInfo[i].ContactInfoType == cit {
	// 		return nil, errors.New("contact info type already exist ")
	// 	}
	// }
	newContactInfo, _ := contactinfo.NewContactInfo(cit, civ)
	c.MyContactInfo = append(c.MyContactInfo, newContactInfo)
	return newContactInfo, nil
}

// func (c *Contact) UpdateContactInfo(contactname, field, value string, cit, civ string) (*Contact, error) {

// 	contactUp, err := FindContactID(c.MyContactInfo, contactname)
// 	if err != nil {
// 		return nil, errors.New("Contact does not exist")
// 	}

// 	switch field {
// 	case "cit":
// 		contactUp.MyContactInfo.cit = cit
// 		return contactUp, nil
// 	case "civ":
// 		contactUp.MyContactInfo.civ = civ
// 		return contactUp, nil
// 	default:
// 		return nil, errors.New("Invalid field")
// 	}
// }

func (c *Contact) UpdateContactInfo(contactname, cit, civ string) (*contactinfo.ContactInfo, error) {
	contactUpdate, contactExist := FindContactInfo(c.MyContactInfo, cit)
	if !contactExist {
		return nil, errors.New("contact do not exist")
	}
	contactUpdate.UpdateContactInfo(cit, civ)
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", c.MyContactInfo)
	return contactUpdate, nil

}

func FindContactByName(contacts []contactinfo.ContactInfo, name string) (*contactinfo.ContactInfo, error) {
	/*
		for i := 0; i < len(contacts); i++ {
			if contacts[i].contactInfoType == name {
				return &contacts[i], nil
			}
		}
		return nil, errors.New("Contact not found")
	*/
	result, err := contactinfo.FindContactInfo(contacts, name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindContactInfo(contactSlice []*contactinfo.ContactInfo, cit string) (*contactinfo.ContactInfo, bool) {
	for i := 0; i < len(contactSlice); i++ {
		if contactSlice[i].ContactInfoType == cit {
			return contactSlice[i], true
		}
	}
	return nil, false
}

func (c *Contact) DeleteContactInfo(contactinfotype string, concontactname string) {

	for i := 0; i < len(c.MyContactInfo); i++ {
		if c.MyContactInfo[i].ContactInfoType == contactinfotype {
			c.MyContactInfo = append(c.MyContactInfo[:i], c.MyContactInfo[i+1:]...)
		}
	}
	//return , nil
}
