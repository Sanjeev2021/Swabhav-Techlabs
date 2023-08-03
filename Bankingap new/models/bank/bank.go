package bank

import (
	"github.com/jinzhu/gorm"
)

type Bank struct {
	gorm.Model
	//	AdminId uint `json:"AdminId"`
	//User   User `json:"User"`
	BankName   string `json:"BankName" gorm:"type:varchar(100)"`
	AdminRefer uint   `json:"AdminRefer"`
}
