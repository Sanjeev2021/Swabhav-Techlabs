package contactinfo

//CRUD OPERATIONS IN CONTACTINFO

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// type of contact
type ContactInfo struct {
	ID               uuid.UUID
	ContactInfoType  string
	ContactInfoValue string
}

func NewContactInfo(contactinfotype, contactinfovalue string) (*ContactInfo, error) {

	return &ContactInfo{
		ID:               uuid.NewV4(),
		ContactInfoType:  contactinfotype,
		ContactInfoValue: contactinfovalue,
	}, nil

}

//func ReadNewContact() ([]ContactInfo , error) {

//}

// func (c *ContactInfo) CreateContactInfo(contactinfotype string) (*ContactInfo, error) {
// 	if !c.IsAdmin {
// 		return nil, errors.New(c.FirstName + "not admin")
// 	}
// 	_, contactExist := FindContactInfo(c.Mycontacts, contactinfotype)
// 	if !contactExist {
// 		return nil, errors.New("Contact already exists")
// 	}

// 	c.Mycontacts = append(c.Mycontacts, newContactInfo)
// 	return newContactInfo, nil
// }

func FindContactInfo(list []ContactInfo, contactName string) (*ContactInfo, error) {
	for i := 0; i < len(list); i++ {
		if list[i].ContactInfoType == contactName {
			return &list[i], nil
		}
	}
	return nil, errors.New("no contact info found")
}

// func UpdateContact(list []ContactInfo, contactName string, field string, value string) (*ContactInfo, error) {
// 	contact, err := FindContactInfo(list, contactName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	switch field {
// 	case "contactInfoType":
// 		contact.ContactInfoType = value
// 		return contact, nil
// 	case "contactInfoValue":
// 		contact.ContactInfoValue = value
// 		return contact, nil
// 	default:
// 		return nil, errors.New("invalid field")
// 	}
// }

// func (c *ContactInfo) DeleteContactInfo(list []ContactInfo, ID uuid.UUID) ([]ContactInfo, error) {
// 	for i := 0; i < len(list); i++ {
// 		if list[i].ID == ID {
// 			return append(list[:i], list[i+1:]...), nil
// 		}

// 	}
// 	return nil, errors.New("USER does not exist")
// }

func (c *ContactInfo) UpdateContactInfo(cit, civ string) *ContactInfo {
	c.ContactInfoType = cit
	c.ContactInfoValue = civ
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>..", c.ContactInfoValue)
	return c
}
