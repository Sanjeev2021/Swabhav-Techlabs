package bank

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"bankingApp/Admin"
	"bankingApp/database"
)

type Bank struct {
	gorm.Model
	ID       uint
	BankName string
	AdminID  uint
	Admin    Admin.Admin
}

func init() {
	db := database.GetDB()
	database.Migrate(db, &Bank{})
}

func CreateBank(bankname string, adminId uint) (*Bank, error) {
	db := database.GetDB()

	Bank := &Bank{
		BankName: bankname,
		AdminID:  adminId,
	}

	result := db.Create(Bank)
	if result.Error != nil {
		return nil, result.Error
	}
	return Bank, nil
}

// GetBank by id
func GetBankById(id uint) (*Bank, error) {
	db := database.GetDB()

	var bank Bank

	result := db.First(&bank, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bank, nil
}

func DeleteBank(id uint) (*Bank, error) {
	db := database.GetDB()

	bank, err := GetBankById(id)
	if err != nil {
		return nil, err
	}

	//db = db.Model(&bank)

	result := db.Delete(bank)
	if result.Error != nil {
		return nil, result.Error
	}
	return bank, nil
}

// update bank
func UpdateBank(id uint, updatebank *Bank) (*Bank, error) {
	db := database.GetDB()

	bank, err := GetBankById(id)
	if err != nil {
		return nil, err
	}

	if updatebank.BankName != "" {
		bank.BankName = updatebank.BankName
	}

	result := db.Save(bank)
	if result.Error != nil {
		return nil, result.Error
	}
	return bank, nil
}

// GetBanks returns all banks
func GetBanks() ([]*Bank, error) {
	db := database.GetDB()

	var banks []*Bank

	result := db.Find(&banks)
	if result.Error != nil {
		return nil, result.Error
	}
	return banks, nil
}

func GetAllBanks(page, pagesize int) ([]*Bank, error) {
	db := database.GetDB()

	var banks []*Bank

	result := db.Offset((page - 1) * pagesize).Limit(pagesize).Find(&banks)
	if result.Error != nil {
		return nil, result.Error
	}
	return banks, nil
}
