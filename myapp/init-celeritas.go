package main

import (
	"log"
	"os"

	"github.com/abhilashdk2016/go-laravel/celeritas"
	"github.com/abhilashdk2016/go-laravel/myapp/handlers"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"

	myHandlers := &handlers.Handlers{
		App: cel,
	}

	cel.InfoLog.Println("Debug is set to", cel.Debug)

	app := &application{
		App:      cel,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
