package main

import (
	"parmod/go-htmx-sample/cmd/router"
	"parmod/go-htmx-sample/infrastructure/server"
	"parmod/go-htmx-sample/internal/application"
)

func main() {

	// dbPool := db.InitDatabase(os.Getenv("DATABASE_URL"))
	app := application.App{
		DbPool: nil,
	}
	appHandler := router.Handler(&app)

	server := server.New(":8080", appHandler)
	server.Start()
}
