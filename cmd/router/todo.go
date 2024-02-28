package router

import (
	"parmod/go-htmx-sample/internal/application"
	"parmod/go-htmx-sample/internal/handlers/home"

	"github.com/gorilla/mux"
)

func todosRouter(app *application.App, router *mux.Router) {
	todosRouter := router.PathPrefix("/todos").Subrouter()

	todosRouter.HandleFunc("/", home.RootHandler(app))
	todosRouter.HandleFunc("/1", home.GetTodos(app))
}
