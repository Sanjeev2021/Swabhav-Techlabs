package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"bankingapp/app"
	"bankingapp/components/log"
	module "bankingapp/modules"
	"bankingapp/repository"
)

func main() {
	// Create New Instace of DB
	log := log.GetLogger(app.SAVELOGSINFILE)
	db := app.NewDBConnection(*log)

	if db == nil {
		log.Print("Db connection failed.")
	}
	defer func() {
		db.Close()
		log.Print("Db closed")
	}()
	var wg sync.WaitGroup
	var repository = repository.NewGormRepository()
	app := app.NewApp("Contact App", db, *log,
		&wg, repository)
	// Initialize app components.
	app.Init()

	module.RegisterModuleRoutes(app, repository)

	// Need to make sure app starts within 60 seconds of deployment so heroku is able to find port.
	go func() {
		err := app.StartServer()
		if err != nil {
			stopApp(app)
		}
	}()
	module.Configure(app)

	// Stop Server On System Call or Interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	stopApp(app)
}
func stopApp(App *app.App) {
	App.Stop()
	App.WG.Wait()
	log.GetLogger(app.SAVELOGSINFILE).Print("App stopped.")
	os.Exit(0)
}
