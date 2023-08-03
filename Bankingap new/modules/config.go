package module

import (
	"bankingapp/app"
	"bankingapp/models/account"
	"bankingapp/models/admin"
	"bankingapp/models/bank"
	"bankingapp/models/user"
)

// Configure will configure all modules
func Configure(appObj *app.App) {
	userModule := user.NewUserModuleConfig(appObj.DB)
	adminModule := admin.NewAdminModuleConfig(appObj.DB)
	bankModule := bank.NewBankModuleConfig(appObj.DB)
	accountModule := account.NewAccountModuleConfig(appObj.DB)

	appObj.MigrateTables([]app.ModuleConfig{userModule})
	appObj.MigrateTables([]app.ModuleConfig{adminModule})
	appObj.MigrateTables([]app.ModuleConfig{bankModule})
	appObj.MigrateTables([]app.ModuleConfig{accountModule})

}
