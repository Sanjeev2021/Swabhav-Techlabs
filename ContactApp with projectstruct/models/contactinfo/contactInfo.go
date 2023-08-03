package contactinfo

import (
	"contactApp/models/contact"

	"github.com/jinzhu/gorm"
)

type ContactInfo struct {
	gorm.Model
	Contact      contact.Contact `gorm:"foreignkey:ContactRefer"`
	ContactRefer uint
	Type         string `json:"Type" gorm:"type:varchar(100)"`
	Value        string `json:"Value" gorm:"type:varchar(100)"`
}
