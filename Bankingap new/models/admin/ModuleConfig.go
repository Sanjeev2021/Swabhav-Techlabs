package admin

import (
	"sync"

	"github.com/jinzhu/gorm"

	"bankingapp/app"
	"bankingapp/components/log"
)

type ModuleConfig struct {
	DB *gorm.DB
}

// NewTestModuleConfig Create New Test Module Config
func NewAdminModuleConfig(db *gorm.DB) *ModuleConfig {
	return &ModuleConfig{
		DB: db,
	}
}

func (config *ModuleConfig) TableMigration(wg *sync.WaitGroup) {
	// Table List
	var models []interface{} = []interface{}{
		&Admin{},
	}
	// Table Migrant
	for _, model := range models {
		if err := config.DB.AutoMigrate(model).Error; err != nil {
			log.GetLogger(app.SAVELOGSINFILE).Print("Auto Migration ==> ", err)
		}
	}
	/*
		if err := config.DB.Model(&Admin{}).
			AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
			log.GetLogger(app.SAVELOGSINFILE).Print("Auto Migration ==> ", err)
		}

	*/
	log.GetLogger(app.SAVELOGSINFILE).Print("Admin Module Configured.")
}
