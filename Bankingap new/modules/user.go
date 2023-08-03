package module

import (
	"bankingapp/app"
	"bankingapp/components/user/controller"
	"bankingapp/components/user/service"
	"bankingapp/repository"
)

func registerUserRoutes(appObj *app.App, repository repository.Repository) {
	defer appObj.WG.Done()
	userService := service.NewUserService(appObj.DB, repository)

	userController := controller.NewUserController(userService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		userController,
	})
}
