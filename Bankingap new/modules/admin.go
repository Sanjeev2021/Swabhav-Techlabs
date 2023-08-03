package module

import (
	"bankingapp/app"
	"bankingapp/components/admin/controller"
	"bankingapp/components/admin/service"
	"bankingapp/repository"
)

func registerAdminRoutes(appObj *app.App, repository repository.Repository) {
	defer appObj.WG.Done()
	adminService := service.NewAdminService(appObj.DB, repository)

	adminController := controller.NewAdminController(adminService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		adminController,
	})
}
