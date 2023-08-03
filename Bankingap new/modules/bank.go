package module

import (
	"bankingapp/app"
	"bankingapp/components/bank/controller"
	"bankingapp/components/bank/service"
	"bankingapp/repository"
)

func registerBankRoutes(appObj *app.App, repository repository.Repository) {
	defer appObj.WG.Done()
	bankService := service.NewBankService(appObj.DB, repository)

	bankController := controller.NewBankController(bankService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		bankController,
	})
}
