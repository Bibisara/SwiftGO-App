package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/bibisara/swiftgo"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init swiftgo
	swi := &swiftgo.SwiftGO{}
	err = swi.New(path)
	if err != nil {
		log.Fatal(err)
	}

	swi.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: swi,
	}

	myHandlers := &handlers.Handlers{
		App: swi,
	}

	app := &application{
		App:        swi,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
