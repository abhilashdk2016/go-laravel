package main

import (
	"github.com/abhilashdk2016/go-laravel/celeritas"
	"github.com/abhilashdk2016/go-laravel/myapp/handlers"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
