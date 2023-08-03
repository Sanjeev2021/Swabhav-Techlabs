package module

import (
	"bankingapp/app"
	"bankingapp/components/account/controller"
	"bankingapp/components/account/service"
	"bankingapp/repository"
)

func registerAccountRoutes(appObj *app.App, repository repository.Repository) {
	defer appObj.WG.Done()
	accountService := service.NewAccountService(appObj.DB, repository)

	accountController := controller.NewAccountController(accountService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		accountController,
	})
}
