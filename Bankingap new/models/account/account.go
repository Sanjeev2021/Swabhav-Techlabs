package account

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	//UserID        uint   `json:"UserID"`
	AccountNumber string `json:"AccountNumber" gorm:"unique;type:varchar(100)"`
	//AccountName   string   `json:"AccountName" gorm:"type:varchar(100)"`
	Balance float64 `json:"Balance" gorm:"type:float"`
	//PassBook      []string `json:"PassBook" gorm:"type:varchar(100)"`
	UserRefer uint `json:"UserRefer"`
	BankRefer uint `json:"BankRefer"`
}
