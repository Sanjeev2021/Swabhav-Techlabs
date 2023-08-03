package account

import (
	"sync"

	"github.com/jinzhu/gorm"

	"bankingapp/app"
	"bankingapp/components/log"
	"bankingapp/models/bank"
)

type ModuleConfig struct {
	DB *gorm.DB
}

// NewTestModuleConfig Create New Test Module Config
func NewAccountModuleConfig(db *gorm.DB) *ModuleConfig {
	return &ModuleConfig{
		DB: db,
	}
}

func (config *ModuleConfig) TableMigration(wg *sync.WaitGroup) {
	// Table List
	var models []interface{} = []interface{}{
		&Account{},
	}
	// Table Migrant
	for _, model := range models {
		if err := config.DB.AutoMigrate(model).Error; err != nil {
			log.GetLogger(app.SAVELOGSINFILE).Print("Auto Migration ==> ", err)
		}
	}
	if err := config.DB.Model(&Account{}).
		AddForeignKey("user_refer", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.GetLogger(app.SAVELOGSINFILE).Print("Auto Migration ==> %s", err)
	}
	if err := config.DB.Model(&bank.Bank{}).
		AddForeignKey("bank_refer", "banks(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.GetLogger(app.SAVELOGSINFILE).Print("Auto Migration ==> %s", err)
	}

	log.GetLogger(app.SAVELOGSINFILE).Print("Test Module Configured.")
}
