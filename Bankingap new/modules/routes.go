package module

import (
	"bankingapp/app"
	"bankingapp/repository"
)

func RegisterModuleRoutes(app *app.App, repository repository.Repository) {
	log := app.Log
	log.Print("============RegisterModuleRoutes.go==============")
	app.WG.Add(4)
	go registerUserRoutes(app, repository)
	go registerAdminRoutes(app, repository)
	go registerBankRoutes(app, repository)
	go registerAccountRoutes(app, repository)

	app.WG.Wait()
}
